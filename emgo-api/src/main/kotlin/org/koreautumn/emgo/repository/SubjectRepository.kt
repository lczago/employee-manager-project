package org.koreautumn.emgo.repository

import org.koreautumn.emgo.model.Subject
import org.springframework.data.jpa.repository.JpaRepository

interface SubjectRepository : JpaRepository<Subject, Long> {
//    fun findByDescriptionAndFieldKnowledge(description: String, fieldKnowledge: FieldKnowledge): Subject
}