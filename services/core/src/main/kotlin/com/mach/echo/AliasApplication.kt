package com.mach.echo

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication


@SpringBootApplication
class AliasApplication

fun main(args: Array<String>) {
    runApplication<AliasApplication>(*args)
}
