package controllers

import (
	"go-rest-service/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	Context *gin.Context
}

type album = dtos.Album

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.998},
}

func (controller *AlbumController) GetAlbums() {
	controller.Context.IndentedJSON(http.StatusOK, albums)
}
