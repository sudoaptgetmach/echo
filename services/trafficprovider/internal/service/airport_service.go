package service

import (
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"sudoaptgetmach.me/trafficprovider/internal/adapter/airportinfo"
	"sudoaptgetmach.me/trafficprovider/internal/domain"
)

type CachedAirport struct {
	Data      domain.EnvironmentMock
	ExpiresAt time.Time
}

var cache struct {
	store map[string]CachedAirport
	mu    sync.RWMutex
}

func InitCache() {
	cache.store = make(map[string]CachedAirport)
}

func decideActiveRunway(runways []airportinfo.RunwaysDto, windDir float64) string {
	bestRunway := "UNK"
	minDiff := 181.0

	for _, r := range runways {
		if r.Closed == "1" {
			continue
		}

		if heading, err := strconv.ParseFloat(r.LeHeadingDegT, 64); err == nil {
			diff := calculateAngularDiff(windDir, heading)
			if diff < minDiff {
				minDiff = diff
				bestRunway = r.LeIdent
			}
		}

		if heading, err := strconv.ParseFloat(r.HeHeadingDegT, 64); err == nil {
			diff := calculateAngularDiff(windDir, heading)
			if diff < minDiff {
				minDiff = diff
				bestRunway = r.HeIdent
			}
		}
	}
	return bestRunway
}

func calculateAngularDiff(a, b float64) float64 {
	diff := math.Abs(a - b)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}

func GetEnvironmentData(icao string) domain.EnvironmentMock {
	cache.mu.RLock()
	item, existe := cache.store[icao]
	cache.mu.RUnlock()
	if existe && item.ExpiresAt.After(time.Now()) {
		return item.Data
	}

	log.Printf("Atualizando cache para %s...", icao)

	var airportRunways = airportinfo.FetchAirportRunways(icao)
	windDir, err := airportinfo.GetAirportWind(icao)
	if err != nil {
		log.Printf("Erro ao pegar o vento de %s: %v", icao, err)
		windDir = 0
	}

	newData := domain.EnvironmentMock{
		ActiveRunway: decideActiveRunway(airportRunways, windDir),
		AssignedSid:  "VATBRZ1A",
		Qnh:          "1013",
	}

	cache.mu.Lock()
	cache.store[icao] = CachedAirport{
		Data:      newData,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	cache.mu.Unlock()

	return newData
}
