package main

import (
	database "assignment2/connection"
	"assignment2/router"
)

var PORT = ":8080"

func main() {
	database.Init()
	router.InitServer().Run(PORT)
}
