package db

import (
	
	"fmt"
	//"net/http"
	"example/basic_api/models"
	//"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"example/basic_api/services"
	//"example/basic_api/controllers"
)
//-----------------------------------------------//
//FIX CONTROLLERS WALA ISSUE FOR CALLING SET DB
//-----------------------------------------------//
type AlbumDB interface {
	AddAlbums(album models.Album) (models.Album, error)
	UpdateAlbum(id string,album models.Album) (models.Album, error)
	DeleteAlbum(id string) (error)
	GetAlbums() ([]models.Album)
	GetAlbumById(id string) (models.Album, error)
}


var DB *services.GormAlbumDB
//nil ptr aa rha constructor bnao and setdb
func ConnectDB(){
	// DB:= services.NewGormAlbumDB()
	dsn := "root:abcd@1234@tcp(127.0.0.1:3306)/albums?charset=utf8mb4&parseTime=True&loc=Local"
	DBconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB= services.NewGormAlbumDB(DBconn)
	//controllers.SetDB(DB)
	if err != nil {
	    panic("Cannot connect to DB")
	} else{
		fmt.Println(DB.DBconn)
		fmt.Println("Connected")
	}
	
	err1:=DB.DBconn.AutoMigrate(&models.Album{})
	if err1!=nil{
		panic(err1)
	} else{
		fmt.Println("Migrated")
	}
}
//SETDB UDAA DIYA
// func SetDB(mockDB models.DB) {
// 	DBconn = mockDB
// }
// func SetDB(mockDB *mocks.DB) {
// 	dbconn = mockDB
// }























	// defer func() {
	// 	dbInstance, _ := dbconn.DB()
	// 	_ = dbInstance.Close()
	// 	fmt.Println("closed")
	// }()

	// result:=db.Create(albums)
