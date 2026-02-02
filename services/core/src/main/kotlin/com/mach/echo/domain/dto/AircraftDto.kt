package com.mach.echo.domain.dto

import com.fasterxml.jackson.annotation.JsonProperty

data class AircraftDto(
    val callsign: String,
    val type: String,
    @JsonProperty("WakeTurbulence") val wakeTurbulence: String,
    val transponder: String
)