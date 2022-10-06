package org.koreautumn.emgo.model

import com.fasterxml.jackson.annotation.JsonManagedReference
import javax.persistence.*

@Entity
@Table(name = "subject")
data class Subject(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    var id: Long? = null,

    var description: String,

    @ManyToOne(fetch = FetchType.EAGER)
    @JsonManagedReference
    var fieldKnowledge: FieldKnowledge
)
