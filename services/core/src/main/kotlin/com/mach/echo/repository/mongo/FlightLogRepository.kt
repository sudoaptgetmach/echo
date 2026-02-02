package com.mach.echo.repository.mongo

import com.mach.echo.domain.document.FlightLog
import org.springframework.data.mongodb.repository.MongoRepository
import org.springframework.stereotype.Repository

@Repository
interface FlightLogRepository : MongoRepository<FlightLog, String>