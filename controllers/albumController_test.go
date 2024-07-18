package controllers_test

import (
	"example/basic_api/controllers"
	"example/basic_api/logger"
	"example/basic_api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


 func TestGetAlbums(t *testing.T) {
	 logger.MyLogger.Info("Testing Getalbums")
 	router := gin.Default()
 	mockDB := new(mocks.MockAlbumDB)
	controllers.SetDB(mockDB)
 	router.GET("/albums", controllers.GetAlbums)
 	req, _ := http.NewRequest("GET", "/albums", nil)
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, http.StatusOK, w.Code)
 	assert.JSONEq(t, `[
 		{
 			"id": "1",
 			"title": "Album1",
 			"artist": "Artist1",
 			"price": 10.99
 		},
 		{
 			"id": "2",
 			"title": "Album2",
 			"artist": "Artist2",
 			"price": 12.99
 		},
 		{
 			"id": "3",
 			"title": "Album3",
 			"artist": "Artist3",
 			"price": 15.99
 		}
 	]`, w.Body.String())
	 logger.MyLogger.Info("Test Completed")
 }



 func TestGetAlbumById(t *testing.T) {
	 logger.MyLogger.Info("Testing GetalbumById")
 	mockDB := new(mocks.MockAlbumDB)
 	controllers.SetDB(mockDB)
 	router := gin.Default()
 	router.GET("/albums/:id", controllers.GetAlbumById)
 	req, _ := http.NewRequest("GET","/albums/1" , nil)
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, http.StatusOK, w.Code)
 	assert.JSONEq(t, `{"id":"1", "title":"Mocked Album", "artist":"Mocked Artist", "price":19.99}`, w.Body.String())
	 logger.MyLogger.Info("Test Completed")
 }


 func TestAddAlbums(t *testing.T) {
	 logger.MyLogger.Info("Testing AddAlbums")
 	mockDB := new(mocks.MockAlbumDB)
 	controllers.SetDB(mockDB)
 	router := gin.Default()
 	router.POST("/albums", controllers.AddAlbums)
 	reqBody := `{"id":"4","title":"Album4","artist":"Artist4","price":29.99}`
 	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(reqBody))
 	req.Header.Set("Content-Type", "application/json")
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, http.StatusCreated, w.Code)
 	assert.JSONEq(t, reqBody, w.Body.String())
	 logger.MyLogger.Info("Test Completed")
 }

 func TestAddAlbumsEmpty(t *testing.T) {
	 logger.MyLogger.Info("Testing AddAlbums with empty input")
 	mockDB := new(mocks.MockAlbumDB)
 	controllers.SetDB(mockDB)
 	router := gin.Default()
 	router.POST("/albums", controllers.AddAlbums)
 	reqBody := ``
 	req, _ := http.NewRequest("POST", "/albums", strings.NewReader(reqBody))
 	req.Header.Set("Content-Type", "application/json")
 	w := httptest.NewRecorder()
 	router.ServeHTTP(w, req)
 	assert.Equal(t, 400, w.Code)
	 logger.MyLogger.Info("Test Completed")
 }

 func TestUpdateAlbum(t *testing.T){
	 logger.MyLogger.Info("Testing UpdateAlbum")
	mockDB:=new(mocks.MockAlbumDB)
	controllers.SetDB(mockDB)
	router:=gin.Default()
	router.PUT("/albums/:id",controllers.UpdateAlbum)
	reqBody := `{"id":"1","title":"Updated Album 1","artist":"Updated Artist 1","price":29.99}`
	req,_:=http.NewRequest("PUT","/albums/1",strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w:=httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, reqBody, w.Body.String())
	 logger.MyLogger.Info("Test Completed")
 }

 func TestDeleteAlbum(t *testing.T){
	 logger.MyLogger.Info("Testing DeleteAlbum")
	mockDB:=new(mocks.MockAlbumDB)
	controllers.SetDB(mockDB)
	router:=gin.Default()
	router.DELETE("/albums/:id",controllers.DeleteAlbum)
	req,_:=http.NewRequest("DELETE","/albums/1",nil)
	w:=httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "album deleted successfully"}`, w.Body.String())
	 logger.MyLogger.Info("Test Completed")
 }