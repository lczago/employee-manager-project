package org.koreautumn.emgo.controller

import org.koreautumn.emgo.dto.NewSubjectForm
import org.koreautumn.emgo.model.Subject
import org.koreautumn.emgo.service.SubjectService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import org.springframework.web.util.UriComponentsBuilder
import javax.transaction.Transactional
import javax.validation.Valid

@RestController
@RequestMapping("/subject")
class SubjectController(
    private val service: SubjectService
) {
    @GetMapping
    fun listAll(): List<Subject> {
        return service.listAll()
    }

    @PostMapping
    @Transactional
    fun createSubject(
        @RequestBody @Valid subjectBody: NewSubjectForm,
        uriBuilder: UriComponentsBuilder
    ): ResponseEntity<Subject> {
        val subject = service.createSubject(subjectBody)
        val uri = uriBuilder.path("/subject/${subject.id}").build().toUri()
        return ResponseEntity.created(uri).body(subject)
    }

    @PutMapping
    @Transactional
    fun editSubject(@RequestBody @Valid subjectBody: Subject): ResponseEntity<Subject> {
        val subject = service.updateSubject(subjectBody)
        return ResponseEntity.ok(subject)
    }

    @DeleteMapping("/{id}")
    @Transactional
    @ResponseStatus(HttpStatus.NO_CONTENT)
    fun deleteSubject(@PathVariable id: Long) {
        service.deleteSubject(id)
    }

}