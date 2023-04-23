package main

import (
	configure "go-api/config"
	"go-api/database"
	"go-api/routes"
	"log"
)

func main() {
	database.ConnectDB()
	defer database.CloseMongo()

	api := configure.ConfigureApi()
	routes.HandleRequests(api)

	if err := api.Listen(":3000"); err != nil {
		log.Print(err.Error())
	}
}
