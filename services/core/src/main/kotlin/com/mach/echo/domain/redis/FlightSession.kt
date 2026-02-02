package com.mach.echo.domain.redis

import org.springframework.data.annotation.Id
import org.springframework.data.redis.core.RedisHash
import java.io.Serializable

@RedisHash("sessions", timeToLive = 300)
data class FlightSession(
    @Id
    val callsign: String,

    var gameState: String = "CLEARANCE_REQ",
    var score: Int = 0,

    var activeRunway: String,
    var assignedSid: String,
    var routeString: String,
    var lastMetar: String,

    var lastSeen: Long = System.currentTimeMillis()
) : Serializable