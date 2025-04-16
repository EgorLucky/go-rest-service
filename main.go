package main

import (
	"go-rest-service/controllers"

	"github.com/gin-gonic/gin"
)

type albumController = controllers.AlbumController

func main() {
	router := gin.Default()
	router.GET("/albums", func(c *gin.Context) {
		var controller = albumController{Context: c}
		controller.GetAlbums()
	})

	router.Run("localhost:8080")
}
