package main

import (
	"github.com/api-market-data/pkg/router"
)

func main() {
	router := router.Start()
	router.Run(":8080")
}
