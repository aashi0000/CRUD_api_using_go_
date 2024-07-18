package controllers

import (
	"example/basic_api/db"
	"example/basic_api/logger"
	"example/basic_api/models"
	//"example/basic_api/services"
	"net/http"
	"github.com/gin-gonic/gin"
)
var MYalbumDB db.AlbumDB 
func SetDB(temp db.AlbumDB) {
		MYalbumDB = temp
}

func GetAlbums(c *gin.Context) {
	//MYalbumDB=db.DB
	log:=logger.NewLogger()
	albums:=MYalbumDB.GetAlbums()
	log.Info("Albums found and returned")
	c.IndentedJSON(http.StatusOK, albums)
	log.Info("Albums sent")
}

func AddAlbums(c *gin.Context) {
	//MYalbumDB=db.DB
	log:=logger.NewLogger()
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		log.Error("Album not in proper format")
		return
	}
	
	//result:=services.AddAlbums(newalbum)
	result,err:=MYalbumDB.AddAlbums(newalbum)
	if err!=nil{
		log.Error("Album not created")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not created"})
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
	log.Info("Album created and sent")
}

func GetAlbumById(c *gin.Context) {
	//MYalbumDB=db.DB
	log:=logger.NewLogger()
	id := c.Param("id")
	album,err:=MYalbumDB.GetAlbumById(id)
	if err!=nil{
		log.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	log.Info("Album found and sent")
	c.IndentedJSON(http.StatusOK, album)
}
func UpdateAlbum(c *gin.Context) {
	//MYalbumDB=db.DB
	log:=logger.NewLogger()
	id := c.Param("id")
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		log.Error("Album not in correct format")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result,err:=MYalbumDB.UpdateAlbum(id,newalbum)
	if err!=nil{
		log.Error("Album not updated")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not updated"})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
	log.Info("Album found and sent")
}

func DeleteAlbum(c *gin.Context) {
	//MYalbumDB=db.DB
	log:=logger.NewLogger()
	id := c.Param("id")
	err:=MYalbumDB.DeleteAlbum(id)
	if err!=nil{
		log.Error("Album not found")
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "album deleted successfully"} )
	log.Info("Album found and deleted")
}