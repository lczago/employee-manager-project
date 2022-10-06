package org.koreautumn.emgo.mapper

import org.koreautumn.emgo.dto.NewSubjectForm
import org.koreautumn.emgo.model.Subject
import org.koreautumn.emgo.service.FieldKnowledgeService
import org.springframework.stereotype.Component

@Component
class SubjectFormMapper(
    private val fieldKnowledgeService: FieldKnowledgeService
) : Mapper<NewSubjectForm, Subject> {
    override fun map(t: NewSubjectForm): Subject {
        return Subject(
            description = t.description,
            fieldKnowledge = fieldKnowledgeService.searchById(t.fieldKnowledge)
        )
    }

}