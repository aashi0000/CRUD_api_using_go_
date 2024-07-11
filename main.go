package main

import (
	"example/basic_api/db"
	"github.com/gin-gonic/gin"
)


func main() {
	db.ConnectDB()
	router := gin.Default()
	router.GET("/albums", db.GetAlbums)
	router.GET("/albums/:id", db.GetAlbumById)
	router.PUT("/albums/:id", db.UpdateAlbum)
	router.DELETE("/albums/:id", db.DeleteAlbum)
	router.POST("/albums", db.AddAlbums)
	router.Run("localhost:5000")
}

