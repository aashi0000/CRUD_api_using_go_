package controllers
import (
	
	"fmt"
	"net/http"
	"example/basic_api/models"
	"github.com/gin-gonic/gin"
	"example/basic_api/db"
)
func GetAlbums(c *gin.Context) {
	var albums []models.Album
	db.DBconn.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbums(c *gin.Context) {
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		return
	}
	result:=db.DBconn.Create(&newalbum)
	if result.Error!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newalbum)
}

func GetAlbumById(c *gin.Context) {
	var album models.Album
	id := c.Param("id")
	result:=db.DBconn.First(&album,id)
	if result.Error!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var newalbum models.Album
	if err := c.BindJSON(&newalbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(newalbum)
    result:=db.DBconn.Model(&models.Album{}).Where("id = ?",id).Updates(models.Album{ID: id, Title: newalbum.Title, Artist: newalbum.Artist, Price: newalbum.Price})
	//result:=dbconn.Model(&Album{}).Where("id = ?", id).Updates(newalbum)
	if result.Error!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, newalbum)
}

func DeleteAlbum(c *gin.Context) {

	id := c.Param("id")
	if db.DBconn == nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "database connection is nil"})
		return
	}
	result:=db.DBconn.Where("id = ?",id).Delete(&models.Album{})

	//dbconn.Where("id = ?",id).Delete(&Album{})
	if result.Error!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "album deleted successfully"} )
}