package main

import (
	"github.com/api-prices/pkg/router"
)

func main() {
	router := router.Start()
	router.Run(":8080")
}
