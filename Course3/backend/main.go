package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/routes"
)

func main() {
	db.ConnectDB()
	fmt.Println("Server is running in port 0666")
	routes.HandleRequest()
}
