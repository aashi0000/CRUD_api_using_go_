package controllers_test

 import (
 	"fmt"
 	"net/http"
 	"net/http/httptest"
 	"strings"
 	"testing"
 	"gorm.io/gorm"
 	"example/basic_api/db"
 	"example/basic_api/controllers"
 	"example/basic_api/mocks1"
 	"example/basic_api/models"
 	//"example/basic_api/logger"
 	"github.com/gin-gonic/gin"
 	"github.com/stretchr/testify/assert"
 	"github.com/stretchr/testify/mock"
 )

 func TestGetAlbums(t *testing.T) {
 	router := gin.Default()
 	mockDB := new(mocks.DB)
 	albums := []models.Album{	
     {
         ID: "1",
         Title: "Blue Train",
         Artist: "John Coltrane",
         Price: 56.99,
     },
     {
         ID: "2",
         Title: "Jeru",
         Artist: "Gerry Mulligan",
         Price: 17.99,
     },
     {
         ID: "3",
        Title: "Clifford Brownn",
         Artist: "Sarah",
         Price: 15.99,
     },
 	}
 	mockDB.On("Find", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
 		out := args.Get(0).(*[]models.Album)
 		*out = albums
 	})
 	db.SetDB(mockDB)
 	router.GET("/albums", controllers.GetAlbums)
 	req, _ := http.NewRequest("GET", "/albums", nil)
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, http.StatusOK, w.Code)
 	assert.JSONEq(t, `[
 		{
 			"id": "1",
 			"title": "Blue Train",
 			"artist": "John Coltrane",
 			"price": 56.99
 		},
 		{
 			"id": "2",
 			"title": "Jeru",
 			"artist": "Gerry Mulligan",
 			"price": 17.99
 		},
 		{
 			"id": "3",
 			"title": "Clifford Brownn",
 			"artist": "Sarah",
 			"price": 15.99
 		}
 	]`, w.Body.String())
 }



 func TestGetAlbumById(t *testing.T) {
 	mockDB := new(mocks.DB)
 	db.SetDB(mockDB)
 	exp_album := models.Album{
         ID: "1",
         Title: "Blue Train",
         Artist: "John Coltrane",
         Price: 56.99,
     }
 	mockDB.On("First", mock.AnythingOfType("*models.Album"), "1").Run(func(args mock.Arguments) {
 		album := args.Get(0).(*models.Album)
 		*album = exp_album
 	}).Return(&gorm.DB{})
 	router := gin.Default()
 	router.GET("/albums/:id", controllers.GetAlbumById)
 	req, _ := http.NewRequest("GET","/albums/1" , nil)
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	fmt.Println("req: ",w.Code)
 	assert.Equal(t, http.StatusOK, w.Code)
 	assert.JSONEq(t, `{"id":"1", "title":"Blue Train", "artist":"John Coltrane", "price":56.99}`, w.Body.String())
 }


 func TestAddAlbums(t *testing.T) {
	
 	mockDB := new(mocks.DB)
 	db.SetDB(mockDB)
 	newAlbum := models.Album{ID: "4", Title: "Album4", Artist: "Artist4", Price: 29.99}

 	mockDB.On("Create", mock.AnythingOfType("*models.Album")).Run(func(args mock.Arguments) {
 		result := args.Get(0).(*models.Album)
 		*result = newAlbum
 	}).Return(&gorm.DB{})

	
 	router := gin.Default()
 	router.POST("/albums", controllers.AddAlbums)

 	reqBody := `{"id":"4","title":"Album4","artist":"Artist4","price":29.99}`
 	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(reqBody))
 	req.Header.Set("Content-Type", "application/json")

 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)

 	assert.Equal(t, http.StatusCreated, w.Code)
 	assert.JSONEq(t, reqBody, w.Body.String())
 }
 func TestAddAlbumsEmpty(t *testing.T) {
	
 	mockDB := new(mocks.DB)
 	db.SetDB(mockDB)
 	mockDB.On("Create", mock.AnythingOfType("*models.Album")).Return(&gorm.DB{})

 	router := gin.Default()
 	router.POST("/albums", controllers.AddAlbums)
 	reqBody := ``
 	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(reqBody))
 	req.Header.Set("Content-Type", "application/json")
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, 400, w.Code)
 }









