package main

import (
	"fmt"
	"net/http"
	"todolist-go/routes"
)

func main() {
	fmt.Println("Server is running in port 0666")
	routes.LoadRoutes()
	http.ListenAndServe(":0666", nil)
}
