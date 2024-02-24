package main

import (
	"github.com/api-prices/pkg/router"
)

func main() {
	// Initialize routes
	r := router.SetupRouter()

	// Run the server
	r.Run(":8080")
}
