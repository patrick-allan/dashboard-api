package main

import (
	"fmt"

	"dashboard-api/pkg/config"
	"dashboard-api/pkg/handlers"
)

func main() {
	fmt.Println("Dashboard - API")
	config.Load()

	handlers.StartHTTPServer()
}
