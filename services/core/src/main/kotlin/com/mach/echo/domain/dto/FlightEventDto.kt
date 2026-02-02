package com.mach.echo.domain.dto

import com.fasterxml.jackson.annotation.JsonProperty

data class FlightEventDto(
    @JsonProperty("scenario_id") val scenarioId: String,
    val source: String,
    val aircraft: AircraftDto,
    @JsonProperty("flight_plan") val flightPlan: FlightPlanDto,
    @JsonProperty("environment_mock") val environment: EnvironmentDto,
    @JsonProperty("expected_state") val expectedState: String
)