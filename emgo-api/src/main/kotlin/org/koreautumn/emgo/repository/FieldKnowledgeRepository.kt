package org.koreautumn.emgo.repository

import org.koreautumn.emgo.model.FieldKnowledge
import org.koreautumn.emgo.model.Subject
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.data.jpa.repository.Query
import org.springframework.data.repository.query.Param

interface FieldKnowledgeRepository : JpaRepository<FieldKnowledge, Long> {
    fun findAllByDescription(description: String): List<FieldKnowledge>

    @Query("select s from Subject s where s.fieldKnowledge.id = :pId ")
    fun listAllSubjects(@Param("pId") id: Long): List<Subject>
}