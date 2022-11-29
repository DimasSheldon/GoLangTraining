package main

import (
	"assignment3/router"
)

func main() {
	router.StartServer().Run(":8080")
}
