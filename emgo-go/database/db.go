package database

import (
	"emgo-go/models"
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
	&models.FieldKnowledge{},
	&models.Subject{},
}

func ConnectDB() {
	//connectionString := "host=localhost user=postgres password=pteste123 dbname=project_emgo port=5432 sslmode=disable"
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
