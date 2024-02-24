package router

import (
	"github.com/api-prices/pkg/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	cryptoController := &controller.CryptocurrencyController{}

	router.GET("/cryptocurrencies", cryptoController.GetCryptocurrencies)

	return router
}
