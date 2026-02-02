package com.mach.echo.domain.airport

import jakarta.persistence.Column
import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table

@Entity
@Table(name = "airports")
data class Airport(
    @Id
    @Column(length = 4)
    val icao: String,

    val name: String,

    @Column(name = "elevation_ft")
    val elevation: Int
)
