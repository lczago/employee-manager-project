package database

import (
	"go-api/database/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

var structModels = []interface{}{
	&entities.FieldKnowledgeEntity{},
	&entities.SubjectEntity{},
}

func ConnectDB() {
	DB, err = gorm.Open(sqlite.Open("project_emgo.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Panic("Error while trying to connect to the database!")
	}

	err := DB.AutoMigrate(structModels...)
	if err != nil {
		panic(err.Error())
	}
}
