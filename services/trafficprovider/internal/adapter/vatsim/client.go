package vatsim

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"sudoaptgetmach.me/trafficprovider/internal/domain"
)

type Data struct {
	Cid            int     `json:"cid"`
	Name           string  `json:"name"`
	Callsign       string  `json:"callsign"`
	Server         string  `json:"server"`
	PilotRating    int     `json:"pilot_rating"`
	MilitaryRating int     `json:"military_rating"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	Altitude       int     `json:"altitude"`
	Groundspeed    int     `json:"groundspeed"`
	Transponder    string  `json:"transponder"`
	Heading        int     `json:"heading"`
	QnhIHg         float64 `json:"qnh_i_hg"`
	QnhMb          int     `json:"qnh_mb"`
	FlightPlan     struct {
		FlightRules         string `json:"flight_rules"`
		Aircraft            string `json:"aircraft"`
		AircraftFaa         string `json:"aircraft_faa"`
		AircraftShort       string `json:"aircraft_short"`
		Departure           string `json:"departure"`
		Arrival             string `json:"arrival"`
		Alternate           string `json:"alternate"`
		CruiseTas           string `json:"cruise_tas"`
		Altitude            string `json:"altitude"`
		Deptime             string `json:"deptime"`
		EnrouteTime         string `json:"enroute_time"`
		FuelTime            string `json:"fuel_time"`
		Remarks             string `json:"remarks"`
		Route               string `json:"route"`
		RevisionId          int    `json:"revision_id"`
		AssignedTransponder string `json:"assigned_transponder"`
	} `json:"flight_plan"`
	LogonTime   time.Time `json:"logon_time"`
	LastUpdated time.Time `json:"last_updated"`
}

type Response struct {
	Pilots []PilotDTO `json:"pilots,omitempty"`
}

type PilotDTO struct {
	Callsign      string         `json:"callsign"`
	FlightPlanDTO *FlightPlanDTO `json:"flight_plan"`
	LastUpdated   time.Time      `json:"last_updated"`
}

type FlightPlanDTO struct {
	FlightRules         string `json:"flight_rules"`
	Aircraft            string `json:"aircraft"`
	AircraftShort       string `json:"aircraft_short"`
	Departure           string `json:"departure"`
	Arrival             string `json:"arrival"`
	Alternate           string `json:"alternate"`
	CruiseTas           string `json:"cruise_tas"`
	Altitude            string `json:"altitude"`
	Remarks             string `json:"remarks"`
	Route               string `json:"route"`
	AssignedTransponder string `json:"assigned_transponder"`
}

func parseWakeTurbulence(rawAircraftString string) string {
	if rawAircraftString == "" {
		return "UNK"
	}

	parts := strings.Split(rawAircraftString, "/")
	if len(parts) < 2 {
		return "UNK"
	}

	wakeTurbulence := parts[1]
	if len(wakeTurbulence) == 0 {
		return "UNK"
	}

	wakeTurbChar := string(wakeTurbulence[0])

	switch wakeTurbChar {
	case "L":
		return "LIGHT"
	case "M":
		return "MEDIUM"
	case "H":
		return "HEAVY"
	case "J":
		return "SUPER"
	default:
		return "UNK"
	}
}

func isRelevantFlight(fp *FlightPlanDTO) bool {
	if fp == nil {
		return false
	}

	prefixes := []string{"SB", "SD", "SI", "SJ", "SN", "SS", "SW"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(fp.Departure, prefix) || strings.HasPrefix(fp.Arrival, prefix) {
			return true
		}
	}

	return false
}

func FetchData() []domain.Flight {
	resp, err := http.Get("https://data.vatsim.net/v3/vatsim-data.json")
	if err != nil {
		log.Fatalln(err)
	}

	var data Response

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
	}

	flights := make([]domain.Flight, 0)

	for _, p := range data.Pilots {

		if !isRelevantFlight(p.FlightPlanDTO) {
			continue
		}

		var dFlightPlan domain.FlightPlan

		if p.FlightPlanDTO != nil {
			dFlightPlan = domain.FlightPlan{
				Departure:   p.FlightPlanDTO.Departure,
				Arrival:     p.FlightPlanDTO.Arrival,
				RouteString: p.FlightPlanDTO.Route,
			}
		}

		newFlight := domain.Flight{
			ScenarioId: uuid.New().String(),
			Source:     domain.VatsimLive,

			Aircraft: domain.Aircraft{
				Callsign:       p.Callsign,
				Type:           p.FlightPlanDTO.AircraftShort,
				WakeTurbulence: parseWakeTurbulence(p.FlightPlanDTO.Aircraft),
				Transponder:    p.FlightPlanDTO.AssignedTransponder,
			},

			FlightPlan: dFlightPlan,

			EnvironmentMock: domain.EnvironmentMock{
				ActiveRunway: "10L",
				AssignedSid:  "ESORU1A",
				Qnh:          "1013",
			},

			ExpectedState: domain.Clearance,
		}

		flights = append(flights, newFlight)
	}

	return flights
}
