package com.mach.echo.domain.role

import com.mach.echo.domain.enums.RoleName
import jakarta.persistence.*
import org.hibernate.annotations.NaturalId

@Entity(name = "role")
@Table(name = "roles")
data class Role(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long,

    @Enumerated(EnumType.STRING)
    @NaturalId
    @Column(length = 60)
    val name: RoleName = RoleName.USER
)
