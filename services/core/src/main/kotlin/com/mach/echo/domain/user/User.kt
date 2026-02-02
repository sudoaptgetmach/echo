package com.mach.echo.domain.user

import com.mach.echo.domain.role.Role
import jakarta.persistence.*
import org.hibernate.annotations.CreationTimestamp
import java.time.LocalDateTime

@Table(name = "users")
@Entity(name = "User")
data class User(
    @Id
    var id: Long? = 0L,

    @Column(unique = true, nullable = false)
    val cid: Long = 0L,

    @Column(unique = true, nullable = false)
    val username: String,

    val xp: Long = 0,

    val level: Int = 1,

    @ManyToMany(fetch = FetchType.EAGER)
    @JoinTable(
        name = "user_roles",
        joinColumns = [JoinColumn(name = "user_id")],
        inverseJoinColumns = [JoinColumn(name = "role_id")]
    )
    var roles: MutableList<Role> = mutableListOf(),

    @CreationTimestamp
    @Column(name = "created_at", updatable = false)
    val createdAt: LocalDateTime = LocalDateTime.now()
)