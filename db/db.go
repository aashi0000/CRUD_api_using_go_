package db

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
var dbconn *gorm.DB
var err error
func ConnectDB(){
	//create a database albums first, the modify the connection string to your userid and password
	dsn := "<YOUR-USERID>:<YOUR-PASSWORD>@tcp(127.0.0.1:3306)/albums?charset=utf8mb4&parseTime=True&loc=Local"
	dbconn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	    panic("Cannot connect to DB")
	} else{
		fmt.Println(dbconn)
		fmt.Println("\n\n\n\nConnected\n\n\n\n")
	}
	
	err1:=dbconn.AutoMigrate(&album{})
	if err1!=nil{
		panic(err1)
	} else{
		fmt.Println("\n\n\n\nMigrated\n\n\n\n")
	}
}

func GetAlbums(c *gin.Context) {
	var albums []album
	dbconn.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbums(c *gin.Context) {
	var newalbum album
	if err := c.BindJSON(&newalbum); err != nil {
		return
	}
	result:=dbconn.Create(&newalbum)
	if result.RowsAffected == 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not created"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newalbum)
}

func GetAlbumById(c *gin.Context) {
	var Album album
	id := c.Param("id")
	result:=dbconn.First(&Album,id)
	if result.RowsAffected == 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, Album)
}
func UpdateAlbum(c *gin.Context) {
	id := c.Param("id")
	var newalbum album
	if err := c.BindJSON(&newalbum); err != nil {
		return
	}
	result:=dbconn.Model(&album{}).Where("id = ?",id).Updates(album{ID: id, Title: newalbum.Title, Artist: newalbum.Artist, Price: newalbum.Price})
	if result.RowsAffected == 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album to be updated not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, newalbum)
}

func DeleteAlbum(c *gin.Context) {
	id := c.Param("id")
	result:=dbconn.Where("id = ?",id).Delete(&album{})
	if result.RowsAffected == 0{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album to be deleted not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "album deleted successfully"} )
}





















	// defer func() {
	// 	dbInstance, _ := dbconn.DB()
	// 	_ = dbInstance.Close()
	// 	fmt.Println("closed")
	// }()

	// result:=db.Create(albums)
