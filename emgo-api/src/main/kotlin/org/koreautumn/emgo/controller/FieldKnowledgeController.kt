package org.koreautumn.emgo.controller

import org.koreautumn.emgo.model.FieldKnowledge
import org.koreautumn.emgo.model.Subject
import org.koreautumn.emgo.service.FieldKnowledgeService
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import org.springframework.web.util.UriComponentsBuilder
import javax.transaction.Transactional
import javax.validation.Valid

@RestController
@RequestMapping("/field_knowledge")
class FieldKnowledgeController(
    private val service: FieldKnowledgeService
) {
    @GetMapping
    fun listAll(
        @RequestParam(required = false) description: String?
    ): List<FieldKnowledge> {
        return service.listAll(description)
    }

    @GetMapping("/{id}")
    fun searchFieldKnowledgeById(@PathVariable id: Long): FieldKnowledge {
        return service.searchById(id)
    }

    @GetMapping("/{id}/subjects")
    fun listAllSubjects(@PathVariable id: Long): List<Subject> {
        return service.listAllSubjects(id)
    }

    @PostMapping
    @Transactional
    fun createFieldKnowledge(
        @RequestBody fieldKnowledgeBody: FieldKnowledge,
        uriBuilder: UriComponentsBuilder
    ): ResponseEntity<FieldKnowledge> {
        val fieldKnowledge = service.createFieldKnowledge(fieldKnowledgeBody)
        val uri = uriBuilder.path("/field_knowledge/${fieldKnowledge.id}").build().toUri()
        return ResponseEntity.created(uri).body(fieldKnowledge)
    }

    @PutMapping
    @Transactional
    fun editFieldKnowledge(@RequestBody @Valid fieldKnowledgeBody: FieldKnowledge): ResponseEntity<FieldKnowledge> {
        val fieldKnowledge = service.updateFieldKnowledge(fieldKnowledgeBody)
        return ResponseEntity.ok(fieldKnowledge)
    }

    @DeleteMapping("/{id}")
    @Transactional
    @ResponseStatus(HttpStatus.NO_CONTENT)
    fun deleteFieldKnowledge(@PathVariable id: Long) {
        service.deleteFieldKnowledge(id)
    }
}