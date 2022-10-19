package main

import (
	"github.com/GinAPIRest-go/database"

	"github.com/GinAPIRest-go/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
