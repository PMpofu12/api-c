package handler

import (
	"net/http"

	"github.com/api-prices/pkg/model"
	"github.com/api-prices/pkg/repository"
	"github.com/gin-gonic/gin"
)

// getAlbum responds with the list of all Album as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repository.Albums)
}

// postAlbum adds an album from JSON received in the request body.
func PostAlbum(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	repository.Albums = append(repository.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of Album, looking for
	// an album whose ID value matches the parameter.
	for _, a := range repository.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
