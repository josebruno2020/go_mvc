package main

import (
	"net/http"
	"webapp/routes"
)

func main() {
	routes.SetupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
