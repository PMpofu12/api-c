package router

import (
	"github.com/api-prices/pkg/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {
	Router := gin.Default()

	Router.GET("/albums", handler.GetAlbums)
	Router.GET("/albums/:id", handler.GetAlbumByID)
	Router.POST("/albums", handler.PostAlbum)
	Router.Run("localhost:8080")
}
