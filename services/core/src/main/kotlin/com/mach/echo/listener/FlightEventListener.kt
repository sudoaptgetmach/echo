package com.mach.echo.listener

import com.mach.echo.domain.document.FlightLog
import com.mach.echo.domain.dto.FlightEventDto
import com.mach.echo.repository.mongo.FlightLogRepository
import com.mach.echo.service.GameService
import org.slf4j.LoggerFactory
import org.springframework.amqp.rabbit.annotation.RabbitListener
import org.springframework.stereotype.Component
import tools.jackson.databind.ObjectMapper

@Component
class FlightEventListener(
    private val objectMapper: ObjectMapper,
    private val logRepository: FlightLogRepository,
    private val gameService: GameService
) {
    private val logger = LoggerFactory.getLogger(FlightEventListener::class.java)

    @RabbitListener(queues = ["flight_events"])
    fun handleFlightUpdate(messageRaw: String) {
        try {
            val flightDto = objectMapper.readValue(messageRaw, FlightEventDto::class.java)

            logRepository.save(
                FlightLog(
                    callsign = flightDto.aircraft.callsign,
                    rawPayload = messageRaw,
                    success = true
                )
            )

            logger.info("Voo processado: ${flightDto.aircraft.callsign} na pista ${flightDto.environment.activeRunway}")
            gameService.processFlightUpdate(flightDto)
        } catch (e: Exception) {
            logger.error("Erro ao processar mensagem: ${e.message}")

            logRepository.save(
                FlightLog(
                    callsign = "UNKNOWN",
                    rawPayload = messageRaw,
                    success = false
                )
            )
        }
    }
}