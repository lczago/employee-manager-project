package org.koreautumn.emgo.service

import org.koreautumn.emgo.model.FieldKnowledge
import org.koreautumn.emgo.model.Subject
import org.koreautumn.emgo.repository.FieldKnowledgeRepository
import org.springframework.stereotype.Service

@Service
class FieldKnowledgeService(
    private val repository: FieldKnowledgeRepository
) {
    fun listAll(description: String?): List<FieldKnowledge> {
        if (description != null) {
            return repository.findAllByDescription(description)
        }
        return repository.findAll()
    }

    fun createFieldKnowledge(fieldKnowledge: FieldKnowledge): FieldKnowledge {
        return repository.save(fieldKnowledge)
    }

    fun searchById(id: Long): FieldKnowledge {
        return repository.findById(id).orElse(null)
    }

    fun updateFieldKnowledge(fieldKnowledgeBody: FieldKnowledge): FieldKnowledge? {
        return repository.save(fieldKnowledgeBody)
    }
    fun deleteFieldKnowledge(id: Long) {
        repository.deleteById(id)
    }

    fun listAllSubjects(id: Long): List<Subject> {
        val fieldKnowledge: FieldKnowledge = searchById(id)
        return  fieldKnowledge.subjects
    }
}