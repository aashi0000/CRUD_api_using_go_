package controllers

import (
	"example/basic_api/logger"
	"example/basic_api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

var MYalbumDB models.AlbumDB 
func SetDB(temp models.AlbumDB) {
	MYalbumDB = temp
}

func GetAlbums(c *gin.Context) {
	albums:=MYalbumDB.GetAlbums()
	 logger.MyLogger.Info("Albums found and returned")
	c.IndentedJSON(http.StatusOK, albums)
	 logger.MyLogger.Info("Albums sent")
}

func AddAlbums(c *gin.Context) {
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		 logger.MyLogger.Error("Album not in proper format")
		return
	}
	
	result,err:=MYalbumDB.AddAlbums(newalbum)
	if err!=nil{
		 logger.MyLogger.Error("Album not created")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not created"})
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
	 logger.MyLogger.Info("Album created and sent")
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	album,err:=MYalbumDB.GetAlbumById(id)
	if err!=nil{
		 logger.MyLogger.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	 logger.MyLogger.Info("Album found and sent")
	c.IndentedJSON(http.StatusOK, album)
}
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		 logger.MyLogger.Error("Album not in correct format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result,err:=MYalbumDB.UpdateAlbum(id,newalbum)
	if err!=nil{
		 logger.MyLogger.Error("Album not updated")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not updated"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	 logger.MyLogger.Info("Album found and sent")
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	err:=MYalbumDB.DeleteAlbum(id)
	if err!=nil{
		 logger.MyLogger.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "album deleted successfully"} )
	 logger.MyLogger.Info("Album found and deleted")
}