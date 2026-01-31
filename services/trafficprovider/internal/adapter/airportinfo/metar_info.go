package airportinfo

import _ "time"

type MetarData struct {
	IcaoId     string  `json:"icaoId"`
	ReportTime string  `json:"reportTime"`
	Wdir       float64 `json:"wdir"`
	Wspd       int     `json:"wspd"`
	Altim      int     `json:"altim"`
	RawOb      string  `json:"rawOb"`
}
