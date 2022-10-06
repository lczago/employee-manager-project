create table subject
(
    id                 serial primary key,
    description        varchar(100) not null,
    field_knowledge_id int          not null,
    foreign key (field_knowledge_id) references field_knowledge (id)
);