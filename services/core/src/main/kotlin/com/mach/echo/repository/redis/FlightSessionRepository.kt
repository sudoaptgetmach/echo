package com.mach.echo.repository.redis

import com.mach.echo.domain.redis.FlightSession
import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository

@Repository
interface FlightSessionRepository : CrudRepository<FlightSession, String>