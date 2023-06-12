package main

import (
	router "go-layerd-arc/routes"
)

func main() {
	e := router.Router()
	e.Logger.Fatal(e.Start(":8080"))
}
