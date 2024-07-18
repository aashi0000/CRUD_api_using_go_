package mocks
//not imported ya changed path anywhere
//get up and see code ka uupr wala part tho jo usne bheja
import (
	"github.com/stretchr/testify/mock"
	"example/basic_api/models"
)
type MockAlbumDB struct{
	mock.Mock
}
//createalbum lekin aise hi naam dediya toh ab stick with it lol
func (m *MockAlbumDB) AddAlbums(album models.Album) (models.Album, error) {
	mockAlbum := models.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
	return mockAlbum, nil
}

func (m *MockAlbumDB) UpdateAlbum(id string,album models.Album) (models.Album,error) {
	updatedAlbum := models.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
	return updatedAlbum,nil
	//ya phir inshort return album
}
//removes success msg from delete
func (m *MockAlbumDB) DeleteAlbum(id string) (error) {
	return nil
}
//if it doesnt work then change to find and first (getalbums and getalbumsbyid)
func (m *MockAlbumDB) GetAlbums() ([]models.Album){
	albums := []models.Album{
		{ID: "1", Title: "Album1", Artist: "Artist1", Price: 10.99},
		{ID: "2", Title: "Album2", Artist: "Artist2", Price: 12.99},
		{ID: "3", Title: "Album3", Artist: "Artist3", Price: 15.99},
	}
	return albums
}
//uupr wale se error nhi nikal diya cuz weird error coming in related func
func (m *MockAlbumDB) GetAlbumById(id string) (models.Album, error) {
	mockAlbum := models.Album{
		ID:     id,
		Title:  "Mocked Album",
		Artist: "Mocked Artist",
		Price:  19.99,
	}
	return mockAlbum, nil
}

//now we hav to create an interface as ye standard methods nhi h aur nhi chlenge..toh item services mei karna hog achange
