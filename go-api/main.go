package main

import (
	configure "go-api/config"
	"go-api/database"
	"go-api/routes"
	"log"
)

func main() {
	database.ConnectDB()
	api := configure.ConfigureApi()
	routes.HandleRequests(api)
	err := api.Listen(":8080")
	if err != nil {
		log.Print(err.Error())
	}
}
