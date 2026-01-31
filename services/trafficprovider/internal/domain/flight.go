package domain

import (
	"github.com/google/uuid"
)

type SourceType string
type State string

const (
	VatsimLive SourceType = "VATSIM_LIVE"
	CGNA       SourceType = "CGNA"
)

const (
	Clearance  State = "CLEARANCE_REQ"
	ClearedIfr State = "CLEARED_IFR"
	ClearedVfr State = "CLEARED_VFR"
	TaxiReq    State = "TAXI_REQ"
	Taxi       State = "TAXIING"
	HoldShort  State = "HOLD_SHORT"
	Takeoff    State = "TAKEOFF"
)

type Aircraft struct {
	Callsign    string `json:"callsign,omitempty"`
	Type        string `json:"type,omitempty"`
	Transponder string `json:"transponder,omitempty"`
}

type FlightPlan struct {
	Departure   string `json:"departure,omitempty"`
	Arrival     string `json:"arrival,omitempty"`
	RouteString string `json:"route_string,omitempty"`
}

type EnvironmentMock struct {
	ActiveRunway string `json:"active_runway,omitempty"`
	AssignedSid  string `json:"assigned_sid,omitempty"`
	Qnh          string `json:"qnh,omitempty"`
}

type Flight struct {
	ScenarioId      uuid.UUID       `json:"scenario_id,omitempty"`
	Source          SourceType      `json:"source,omitempty"`
	Aircraft        Aircraft        `json:"aircraft,omitempty"`
	FlightPlan      FlightPlan      `json:"flight_plan,omitempty"`
	EnvironmentMock EnvironmentMock `json:"environment_mock,omitempty"`
	ExpectedState   State           `json:"expected_state,omitempty"`
}
