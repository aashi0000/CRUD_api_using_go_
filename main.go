package main

import (
	"example/basic_api/db"
	"github.com/gin-gonic/gin"
	"example/basic_api/controllers"
)


func main() {
	db.ConnectDB()
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.POST("/albums", controllers.AddAlbums)
	router.Run("localhost:5000")
}

