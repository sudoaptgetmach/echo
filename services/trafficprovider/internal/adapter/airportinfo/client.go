package airportinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type AirportData struct {
	Ident            string    `json:"ident"`
	Type             string    `json:"type"`
	Name             string    `json:"name"`
	LatitudeDeg      float64   `json:"latitude_deg"`
	LongitudeDeg     float64   `json:"longitude_deg"`
	ElevationFt      string    `json:"elevation_ft"`
	Continent        string    `json:"continent"`
	IsoCountry       string    `json:"iso_country"`
	IsoRegion        string    `json:"iso_region"`
	Municipality     string    `json:"municipality"`
	ScheduledService string    `json:"scheduled_service"`
	GpsCode          string    `json:"gps_code"`
	IataCode         string    `json:"iata_code"`
	LocalCode        string    `json:"local_code"`
	HomeLink         string    `json:"home_link"`
	WikipediaLink    string    `json:"wikipedia_link"`
	Keywords         string    `json:"keywords"`
	IcaoCode         string    `json:"icao_code"`
	Runways          []Runways `json:"runways"`
}

type Runways struct {
	Id                     string `json:"id"`
	AirportRef             string `json:"airport_ref"`
	AirportIdent           string `json:"airport_ident"`
	LengthFt               string `json:"length_ft"`
	WidthFt                string `json:"width_ft"`
	Surface                string `json:"surface"`
	Lighted                string `json:"lighted"`
	Closed                 string `json:"closed"`
	LeIdent                string `json:"le_ident"`
	LeLatitudeDeg          string `json:"le_latitude_deg"`
	LeLongitudeDeg         string `json:"le_longitude_deg"`
	LeElevationFt          string `json:"le_elevation_ft"`
	LeHeadingDegT          string `json:"le_heading_degT"`
	LeDisplacedThresholdFt string `json:"le_displaced_threshold_ft"`
	HeIdent                string `json:"he_ident"`
	HeLatitudeDeg          string `json:"he_latitude_deg"`
	HeLongitudeDeg         string `json:"he_longitude_deg"`
	HeElevationFt          string `json:"he_elevation_ft"`
	HeHeadingDegT          string `json:"he_heading_degT"`
	HeDisplacedThresholdFt string `json:"he_displaced_threshold_ft"`
	HeIls                  struct {
		Freq   float64 `json:"freq"`
		Course int     `json:"course"`
	} `json:"he_ils"`
	LeIls struct {
		Freq   float64 `json:"freq"`
		Course int     `json:"course"`
	} `json:"le_ils"`
}

type AirportDataDTO struct {
	Ident    string      `json:"ident"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	IataCode string      `json:"iata_code"`
	Runways  *RunwaysDto `json:"runways"`
}

type RunwaysDto struct {
	Id            string `json:"id"`
	AirportRef    string `json:"airport_ref"`
	AirportIdent  string `json:"airport_ident"`
	Closed        string `json:"closed"`
	LeIdent       string `json:"le_ident"`
	LeHeadingDegT string `json:"le_heading_degT"`
	HeIdent       string `json:"he_ident"`
	HeHeadingDegT string `json:"he_heading_degT"`
}

type Response struct {
	Runways []RunwaysDto `json:"runways,omitempty"`
}

func FetchAirportRunways(icao string) []RunwaysDto {
	resp, err := http.Get(fmt.Sprintf("https://airportdb.io/api/v1/airport/%s?apiToken=%s", icao, os.Getenv("AIRPORTDB_API_KEY")))
	if err != nil {
		log.Println("Erro request:", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Erro:", err)
		return nil
	}

	return data.Runways
}

func GetAirportWind(icao string) (float64, error) {
	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(fmt.Sprintf("https://aviationweather.gov/api/data/metar?ids=%s&format=json&taf=false", icao))
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var data []MetarData
	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, fmt.Errorf("nenhum metar encontrado para %s", icao)
	}

	return data[0].Wdir, nil
}
