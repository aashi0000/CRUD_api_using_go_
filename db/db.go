package db

import (
	"fmt"
	"example/basic_api/config"
	"example/basic_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"example/basic_api/services"
	"example/basic_api/controllers"
	"example/basic_api/logger"
)


var DB *services.GormAlbumDB
func ConnectDB(){
	cfg,err0:=config.LoadConfig("config.json")
	if err0 != nil {
		logger.MyLogger.Error("Failed to load configuration")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,cfg.Database.Password,cfg.Database.Host,cfg.Database.Port,cfg.Database.DBName,)
	DBconn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB= services.NewGormAlbumDB(DBconn)
	controllers.SetDB(DB)
	if err != nil {
		logger.MyLogger.Error("Cannot connect to DB")
	} else{
		logger.MyLogger.Info("CONNECTED")
	}
	
	err1:=DB.DBconn.AutoMigrate(&models.Album{})
	if err1!=nil{
		logger.MyLogger.Error(err1)
	} else{
		logger.MyLogger.Info("MIGRATED")
	}
}























	// defer func() {
	// 	dbInstance, _ := dbconn.DB()
	// 	_ = dbInstance.Close()
	// 	fmt.Println("closed")
	// }()

	// result:=db.Create(albums)
