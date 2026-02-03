package service

import (
	"log"
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

func GetEnvironmentData(icao string) domain.EnvironmentMock {
	cache.mu.RLock()
	item, existe := cache.store[icao]
	cache.mu.RUnlock()
	if existe && item.ExpiresAt.After(time.Now()) {
		return item.Data
	}

	log.Printf("Updating cache for %s...", icao)

	var airportRunways = airportinfo.FetchAirportRunways(icao)
	airportMetar := GetAirportMetar(icao)

	var airportWind = 0
	var windSpeed = 0
	var airportQNH = 1013
	var rawMetar = ""

	if len(airportMetar) > 0 {
		airportWind = airportMetar[0].Wdir
		windSpeed = airportMetar[0].Wspd
		airportQNH = airportMetar[0].Altim
		rawMetar = airportMetar[0].RawOb
	}

	newData := domain.EnvironmentMock{
		ActiveRunway: decideActiveRunway(airportRunways, float64(airportWind)),
		AssignedSid:  "VATBRZ1A",
		Wdir:         airportWind,
		Wspd:         windSpeed,
		Qnh:          airportQNH,
		RawMetar:     rawMetar,
	}

	cache.mu.Lock()
	cache.store[icao] = CachedAirport{
		Data:      newData,
		ExpiresAt: time.Now().Add(30 * time.Minute),
	}
	cache.mu.Unlock()

	return newData
}
