package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"time"

	"sudoaptgetmach.me/trafficprovider/internal/domain"
)

func GetAirportMetar(icao string) []domain.MetarResponseDto {
	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%s&format=json&taf=false", icao))
	if err != nil {
		fmt.Printf("Erro na requisição METAR para %s: %v\n", icao, err)
		return []domain.MetarResponseDto{}
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var data []domain.MetarData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Erro ao decodificar JSON do METAR para %s: %v\n", icao, err)
		return []domain.MetarResponseDto{}
	}

	var metars = make([]domain.MetarResponseDto, 0)

	for _, metar := range data {
		valWdir := toInt(metar.Wdir)
		valWspd := toInt(metar.Wspd)
		valQnh := int(metar.Altim)

		newMetar := domain.MetarResponseDto{
			IcaoId:     metar.IcaoId,
			ReportTime: time.Now(),
			Wdir:       valWdir,
			Wspd:       valWspd,
			Altim:      valQnh,
			RawOb:      metar.RawOb,
		}
		metars = append(metars, newMetar)
	}
	return metars
}

func decideActiveRunway(runways []domain.RunwaysDto, windDir float64) string {
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

func toInt(value interface{}) int {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case float64:
		return int(v)
	case int:
		return v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return int(f)
		}
		return 0
	default:
		return 0
	}
}
