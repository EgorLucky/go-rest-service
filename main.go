package main

import (
	"go-rest-service/controllers"

	"github.com/gin-gonic/gin"
)

type albumController = controllers.AlbumController

func main() {
	router := gin.Default()

	{
		albums := router.Group("/albums")
		{

			albums.GET("", func(ctx *gin.Context) {
				var controller = albumController{Context: ctx}
				controller.GetAlbums()
			})
			albums.POST("", func(ctx *gin.Context) {
				var controller = albumController{Context: ctx}
				controller.PostAlbums()
			})
			albums.GET("/:id", func(ctx *gin.Context) {
				var controller = albumController{Context: ctx}
				controller.GetAlbumByID()
			})
		}

	}
	router.Run("localhost:8080")
}
