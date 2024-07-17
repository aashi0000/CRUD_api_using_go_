package controllers

import (
	
	"example/basic_api/logger"
	"example/basic_api/models"
	"example/basic_api/services"
	
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	log:=logger.NewLogger()
	albums:=services.GetAlbums()
	log.Info("Albums found and returned")
	c.IndentedJSON(http.StatusOK, albums)
	log.Info("Albums sent")
}

func AddAlbums(c *gin.Context) {
	log:=logger.NewLogger()
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		log.Error("Album not in proper format")
		return
	}
	
	result:=services.AddAlbums(newalbum)
	if result!=nil{
		log.Error("Album not created")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not created"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newalbum)
	log.Info("Album created and sent")
}

func GetAlbumById(c *gin.Context) {
	log:=logger.NewLogger()
	id := c.Param("id")
	album,err:=services.GetAlbumById(id)
	if err!=nil{
		log.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	log.Info("Album found and sent")
	c.IndentedJSON(http.StatusOK, album)
}
func UpdateAlbum(c *gin.Context) {
	log:=logger.NewLogger()
	id := c.Param("id")
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		log.Error("Album not in correct format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err:=services.UpdateAlbum(id,newalbum)
	if err!=nil{
		log.Error("Album not updated")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not updated"})
		return
	}
	c.IndentedJSON(http.StatusOK, newalbum)
	log.Info("Album found and sent")
}

func DeleteAlbum(c *gin.Context) {
	log:=logger.NewLogger()
	id := c.Param("id")
	err:=services.DeleteAlbum(id)
	if err!=nil{
		log.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "album deleted successfully"} )
	log.Info("Album found and deleted")
}