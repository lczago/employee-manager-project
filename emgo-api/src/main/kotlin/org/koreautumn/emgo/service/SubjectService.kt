package org.koreautumn.emgo.service

import org.koreautumn.emgo.dto.NewSubjectForm
import org.koreautumn.emgo.mapper.SubjectFormMapper
import org.koreautumn.emgo.model.Subject
import org.koreautumn.emgo.repository.SubjectRepository
import org.springframework.stereotype.Service

@Service
class SubjectService(
    private val repository: SubjectRepository,
    private val subjectFormMapper: SubjectFormMapper
) {
    fun listAll(): List<Subject> {
        return repository.findAll()
    }

    fun createSubject(subjectBody: NewSubjectForm): Subject {
        val subject = subjectFormMapper.map(subjectBody)
        return repository.save(subject)
    }

    fun updateSubject(subject: Subject): Subject? {
        return repository.save(subject)
    }

    fun deleteSubject(id: Long) {
        repository.deleteById(id)
    }
}