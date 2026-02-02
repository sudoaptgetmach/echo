package domain

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
