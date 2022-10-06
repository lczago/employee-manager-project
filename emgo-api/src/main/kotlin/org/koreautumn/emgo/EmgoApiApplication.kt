package org.koreautumn.emgo

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class EmgoApiApplication

fun main(args: Array<String>) {
    runApplication<EmgoApiApplication>(*args)
}
