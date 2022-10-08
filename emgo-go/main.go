package main

import (
	"emgo-go/database"
	"emgo-go/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
