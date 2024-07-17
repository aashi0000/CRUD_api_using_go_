package services

import (
	"example/basic_api/models"
	"example/basic_api/db"
)

func GetAlbums() ([]models.Album) {
	var albums []models.Album
	db.DBconn.Find(&albums)
	return albums
}

func AddAlbums(newalbum models.Album) error {
	result := db.DBconn.Create(&newalbum)
	return result.Error
}

func GetAlbumById(id string) (models.Album, error) {
	var album models.Album
	result := db.DBconn.First(&album, id)
	return album, result.Error
}

func UpdateAlbum(id string, newalbum models.Album) error {
	var album models.Album
	result := db.DBconn.First(&album, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.DBconn.Model(&album).Updates(newalbum)
	return result.Error
}

func DeleteAlbum(id string) error {
	var album models.Album
	result := db.DBconn.First(&album, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.DBconn.Delete(&album)
	return result.Error
}
