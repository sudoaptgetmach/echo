package airportinfo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"sudoaptgetmach.me/trafficprovider/internal/domain"
)

func FetchAirportRunways(icao string) []domain.RunwaysDto {
	resp, err := http.Get(fmt.Sprintf("https://airportdb.io/api/v1/airport/%s?apiToken=%s", icao, os.Getenv("AIRPORTDB_API_KEY")))
	if err != nil {
		log.Println("Error on FetchAirportRunways request:", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	var data domain.Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("Error on AirportInfo Response:", err)
		return nil
	}

	return data.Runways
}
