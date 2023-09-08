package main

import (
	"fmt"

	"github.com/alexsantossilva/go-api-rest/database"
	"github.com/alexsantossilva/go-api-rest/routes"
)

func main() {
	database.ConnectDB()
	fmt.Println("Starting server...")
	routes.HandleResquest()
}
