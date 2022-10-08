package org.koreautumn.emgo.repository

import org.koreautumn.emgo.model.FieldKnowledge
import org.springframework.data.jpa.repository.JpaRepository

interface FieldKnowledgeRepository : JpaRepository<FieldKnowledge, Long> {
    fun findAllByDescription(description: String): List<FieldKnowledge>
}