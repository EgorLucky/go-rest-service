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

func (controller *AlbumController) PostAlbums() {
	context := controller.Context
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func (controller *AlbumController) GetAlbumByID() {
	context := controller.Context
	id := controller.Context.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			context.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
