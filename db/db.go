package db

import (
	
	"fmt"
	//"net/http"
	"example/basic_api/models"
	//"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DBconn models.DB
var err error
func ConnectDB(){
	dsn := "root:abcd@1234@tcp(127.0.0.1:3306)/albums?charset=utf8mb4&parseTime=True&loc=Local"
	DBconn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	    panic("Cannot connect to DB")
	} else{
		fmt.Println(DBconn)
		fmt.Println("Connected")
	}
	
	err1:=DBconn.AutoMigrate(&models.Album{})
	if err1!=nil{
		panic(err1)
	} else{
		fmt.Println("Migrated")
	}
}
func SetDB(mockDB models.DB) {
	DBconn = mockDB
}
// func SetDB(mockDB *mocks.DB) {
// 	dbconn = mockDB
// }























	// defer func() {
	// 	dbInstance, _ := dbconn.DB()
	// 	_ = dbInstance.Close()
	// 	fmt.Println("closed")
	// }()

	// result:=db.Create(albums)
