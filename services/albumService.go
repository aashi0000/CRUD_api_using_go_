package services

import (
	"example/basic_api/models"

	"gorm.io/gorm"
)

type GormAlbumDB struct{
	DBconn *gorm.DB
}
func NewGormAlbumDB(db *gorm.DB) *GormAlbumDB{
	return &GormAlbumDB{DBconn: db}
}

func (db *GormAlbumDB) GetAlbums() ([]models.Album) {
	var albums []models.Album
	db.DBconn.Find(&albums)
	return albums
}

//changesfrom error to albums,err return
func (db *GormAlbumDB) AddAlbums(newalbum models.Album) (models.Album,error) {
	result := db.DBconn.Create(&newalbum)
	return newalbum,result.Error
}

func (db *GormAlbumDB) GetAlbumById(id string) (models.Album, error) {
	var album models.Album
	result := db.DBconn.First(&album, id)
	return album, result.Error
}
//changesfrom error to albums,err return
func (db *GormAlbumDB) UpdateAlbum(id string, newalbum models.Album) (models.Album,error) {
	var album models.Album
	result := db.DBconn.First(&album, id)
	if result.Error != nil {
		return album,result.Error //????
	}
	result = db.DBconn.Model(&album).Updates(newalbum)
	return newalbum,result.Error
}

func (db *GormAlbumDB) DeleteAlbum(id string) error {
	var album models.Album
	result := db.DBconn.First(&album, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.DBconn.Delete(&album)
	return result.Error
}
