package router

import (
	"github.com/api-market-data/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	cryptoController := &controller.CryptocurrencyController{}

	router.GET("/current/market", cryptoController.GetCurrentMarketData)
	router.GET("/current/:cryptocurrency", cryptoController.GetCurrentCoinData)

	// router.GET("/historical/:coin/", cryptoController.GetCryptocurrencies) Pass the date as a paramaeter
	// router.GET("/historical/:coin/", cryptoController.GetCryptocurrencies)

	return router
}
