package main

import (
	"github.com/narozhniyy/test/internal/router"
	"os"
)

func main() {
	// Initialize router with echo framework
	e := router.InitRouter()
	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("API_HOST")))
}
