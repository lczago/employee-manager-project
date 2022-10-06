package org.koreautumn.emgo.mapper

interface Mapper<T, U> {
    fun map(t: T): U
}