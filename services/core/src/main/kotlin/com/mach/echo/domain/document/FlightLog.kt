package com.mach.echo.domain.document

import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document
import java.time.Instant

@Document(collection = "flight_logs")
data class FlightLog(
    @Id val id: String? = null,
    val timestamp: Instant = Instant.now(),
    val callsign: String? = null,
    val success: Boolean = true,
    val rawPayload: String
)