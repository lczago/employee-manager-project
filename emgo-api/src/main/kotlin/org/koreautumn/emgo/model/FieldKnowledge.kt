package org.koreautumn.emgo.model

import com.fasterxml.jackson.annotation.JsonBackReference
import javax.persistence.*
import javax.validation.constraints.NotEmpty

@Entity
@Table(name = "field_knowledge")
data class FieldKnowledge(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long,

    @field:NotEmpty
    var description: String,

    @OneToMany(fetch = FetchType.LAZY, cascade = [CascadeType.REMOVE], mappedBy = "fieldKnowledge")
    @JsonBackReference
    var subjects: List<Subject> = ArrayList()
)
