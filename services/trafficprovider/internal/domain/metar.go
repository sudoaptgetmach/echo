package domain

import "time"

type MetarData struct {
	IcaoId      string      `json:"icaoId"`
	ReceiptTime string      `json:"receiptTime"`
	ObsTime     int         `json:"obsTime"`
	ReportTime  string      `json:"reportTime"`
	Temp        float64     `json:"temp"`
	Dewp        float64     `json:"dewp"`
	Wdir        interface{} `json:"wdir"`
	Wspd        interface{} `json:"wspd"`
	Visib       interface{} `json:"visib"`
	Altim       float64     `json:"altim"`
	QcField     int         `json:"qcField"`
	MetarType   string      `json:"metarType"`
	RawOb       string      `json:"rawOb"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"lon"`
	Elev        int         `json:"elev"`
	Name        string      `json:"name"`
	FltCat      string      `json:"fltCat"`
}

type MetarResponseDto struct {
	IcaoId     string    `json:"icaoId"`
	ReportTime time.Time `json:"reportTime"`
	Wdir       int       `json:"wdir"`
	Wspd       int       `json:"wspd"`
	Altim      int       `json:"altim"`
	RawOb      string    `json:"rawOb"`
}
