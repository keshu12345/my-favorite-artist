package myFavoriteArtistService

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/my-favorite-artist/models"
)

const artistInfo = "https://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=Cher&api_key=2323785c019f7c0509019a8723c35be2&format=json"

// mockery --dir=myFavoriteArtistService --name=MyFavoriteArtistService --output=MyFavoriteArtistService_test/mocks --outpkg=MockMyFavoriteArtistService
type MyFavoriteArtistService interface {
	TopTracks(c *gin.Context, lastFMAPIURL string) (*models.TopTracks, error)
	ArtistInfo(c *gin.Context, lastFMAPIURL string) (*models.ArtistInfo, error)
	ImageInfo(c *gin.Context, lastFMAPIURL string) (*models.ImageWithImageURL, error)
}

// getMyFavoriteArtistService represents the MyFavoriteArtistService for this application
// swagger:model
type getMyFavoriteArtistService struct {
}

func NewMyFavoriteArtistService() MyFavoriteArtistService {
	return getMyFavoriteArtistService{}
}

func (getMyFavoriteArtistService getMyFavoriteArtistService) TopTracks(c *gin.Context,
	lastFMAPIURL string) (*models.TopTracks, error) {
	response, err := http.Get(lastFMAPIURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return &models.TopTracks{}, err
	}

	var topTracks models.TopTracks
	err = json.Unmarshal(body, &topTracks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return &models.TopTracks{}, err
	}

	return &topTracks, nil
}

func (getMyFavoriteArtistService getMyFavoriteArtistService) ArtistInfo(c *gin.Context,
	lastFMAPIURL string) (*models.ArtistInfo, error) {

	response, err := http.Get(lastFMAPIURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return &models.ArtistInfo{}, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return &models.ArtistInfo{}, err
	}

	// Parse the JSON response into the struct
	var lastFMArtist models.ArtistInfo
	err = json.Unmarshal(body, &lastFMArtist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return &models.ArtistInfo{}, err
	}

	return &lastFMArtist, nil
}

func (getMyFavoriteArtistService getMyFavoriteArtistService) ImageInfo(c *gin.Context,
	lastFMAPIURL string) (*models.ImageWithImageURL, error) {

	response, err := http.Get(artistInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return &models.ImageWithImageURL{}, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return &models.ImageWithImageURL{}, err
	}

	// Parse the JSON response into the struct
	var lastFMArtist models.ArtistInfo
	err = json.Unmarshal(body, &lastFMArtist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return &models.ImageWithImageURL{}, err
	}

	var imageInfo models.ImageWithImageURL

	for _, image := range lastFMArtist.Artist.Image {
		imageInfo.Image = append(imageInfo.Image, image)
	}

	imageInfo.URL = lastFMArtist.Artist.URL
	return &imageInfo, nil
}
