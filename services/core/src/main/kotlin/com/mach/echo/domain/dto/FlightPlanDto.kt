package com.mach.echo.domain.dto

import com.fasterxml.jackson.annotation.JsonProperty

data class FlightPlanDto(
    val departure: String,
    val arrival: String,
    @JsonProperty("route_string") val routeString: String,
)
