package myFavoriteArtistService

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// var myFavoriteArtistServiceInterfaceTest MyFavoriteArtistService

func setupRouter() *gin.Engine {
	// Assuming you have a function that sets up the Gin router
	r := gin.Default()
	myFavoriteArtistServiceInterfaceTest := NewMyFavoriteArtistService()
	RegisterMyFavoriteArtistformEndPoint(r, myFavoriteArtistServiceInterfaceTest)
	return r
}

func TestGetTopTracks(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/my-favorite-artist/top-tracks?method=geo.gettoptracks&country=India&api_key=2323785c019f7c0509019a8723c35be2&format=json", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetArtistInfo(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/my-favorite-artist/artist-info?method=artist.getinfo&artist=Cher&api_key=2323785c019f7c0509019a8723c35be2&format=json", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetImageInfo(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/my-favorite-artist/image-info?method=artist.getinfo&artist=Cher&api_key=2323785c019f7c0509019a8723c35be2&format=json", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
