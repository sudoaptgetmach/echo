package com.mach.echo.domain.dto

import com.fasterxml.jackson.annotation.JsonProperty

data class EnvironmentDto(
    @JsonProperty("active_runway") val activeRunway: String,
    @JsonProperty("assigned_sid") val assignedSid: String,
    val wdir: Int,
    val wspd: Int,
    val qnh: Int,
    val rawMetar: String
)