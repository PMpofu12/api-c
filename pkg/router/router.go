package router

import (
	"github.com/api-prices/pkg/handler"

	"github.com/api-prices/pkg/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routing for the API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Apply middleware
	r.Use(handler.LoggingMiddleware)

	// Initialize controllers
	cryptoController := &controller.CryptocurrencyController{}

	// Define routes
	r.GET("/cryptocurrencies", cryptoController.GetCryptocurrencies)

	return r
}
