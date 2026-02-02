package com.mach.echo.service

import com.mach.echo.domain.dto.FlightEventDto
import com.mach.echo.domain.redis.FlightSession
import com.mach.echo.repository.redis.FlightSessionRepository
import org.slf4j.LoggerFactory
import org.springframework.stereotype.Service

@Service
class GameService(
    private val sessionRepository: FlightSessionRepository
) {
    private val logger = LoggerFactory.getLogger(GameService::class.java)

    fun processFlightUpdate(dto: FlightEventDto) {
        val callsign = dto.aircraft.callsign

        val existingSession = sessionRepository.findById(callsign)

        if (existingSession.isPresent) {
            val session = existingSession.get()

            session.activeRunway = dto.environment.activeRunway
            session.assignedSid = dto.environment.assignedSid
            session.lastMetar = dto.environment.rawMetar
            session.lastSeen = System.currentTimeMillis()

            sessionRepository.save(session)
            logger.debug("Sessão atualizada: $callsign (Estado: ${session.gameState})")

        } else {
            val newSession = FlightSession(
                callsign = callsign,
                gameState = "CLEARANCE_REQ",
                activeRunway = dto.environment.activeRunway,
                assignedSid = dto.environment.assignedSid,
                routeString = dto.flightPlan.routeString,
                lastMetar = dto.environment.rawMetar
            )

            sessionRepository.save(newSession)
            logger.info("Nova sessão criada: $callsign na pista ${newSession.activeRunway}")
        }
    }
}