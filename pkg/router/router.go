package router

import (
	"github.com/api-prices/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	cryptoController := &controller.CryptocurrencyController{}

	router.GET("/cryptocurrencies/:currency/", cryptoController.GetCryptocurrencies)

	return router
}
