package org.koreautumn.emgo.dto

import org.jetbrains.annotations.NotNull
import javax.validation.constraints.NotEmpty

data class NewSubjectForm(
    @field:NotEmpty
    val description: String,
    @field:NotNull
    val fieldKnowledge: Long
)
