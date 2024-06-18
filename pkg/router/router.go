package router

import (
	"github.com/api-market-data/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	currentMarketDataController := &controller.CurrentMarketDataController{}

	router.GET("/current/market", currentMarketDataController.GetCurrentMarketData)
	router.GET("/current/coin/:id", currentMarketDataController.GetCurrentCoinData)

	// router.GET("/historical/:coin/", cryptoController.GetCryptocurrencies) Pass the date as a paramaeter
	// router.GET("/historical/:coin/", cryptoController.GetCryptocurrencies)

	return router
}
