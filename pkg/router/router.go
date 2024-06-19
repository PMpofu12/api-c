package router

import (
	"github.com/api-market-data/pkg/controller"

	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	router := gin.Default()

	currentMarketDataController := &controller.CurrentMarketDataController{}
	historicalCoinDataController := controller.HistoricalCoinDataController{}

	router.GET("/current/market", currentMarketDataController.GetCurrentMarketData)
	router.GET("/current/coin/:id", currentMarketDataController.GetCurrentCoinData)
	router.GET("/historical/coin/:id", historicalCoinDataController.GetHistoricalCoinData)

	return router
}
