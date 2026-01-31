package domain

type SourceType string
type State string
type AircraftType string

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

const (
	LIGHT  string = "LIGHT"
	MEDIUM string = "MEDIUM"
	HEAVY  string = "HEAVY"
	SUPER  string = "SUPER"
)

type Aircraft struct {
	Callsign       string `json:"callsign,omitempty"`
	Type           string `json:"type,omitempty"`
	WakeTurbulence string
	Transponder    string `json:"transponder,omitempty"`
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
	ScenarioId      string          `json:"scenario_id,omitempty"`
	Source          SourceType      `json:"source,omitempty"`
	Aircraft        Aircraft        `json:"aircraft,omitempty"`
	FlightPlan      FlightPlan      `json:"flight_plan,omitempty"`
	EnvironmentMock EnvironmentMock `json:"environment_mock,omitempty"`
	ExpectedState   State           `json:"expected_state,omitempty"`
}
