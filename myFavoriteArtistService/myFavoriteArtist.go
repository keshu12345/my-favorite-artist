package myFavoriteArtistService

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/my-favorite-artist/docs"

	// logger "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var fmBaseURL string = ""

var myFavoriteArtistServiceInterface MyFavoriteArtistService

// Endpoint
func RegisterMyFavoriteArtistformEndPoint(g *gin.Engine, mfas MyFavoriteArtistService) {
	fmBaseURL = "https://ws.audioscrobbler.com/2.0/"
	docs.SwaggerInfo.BasePath = "/my-favorite-artist"
	myFavoriteArtist := g.Group("/my-favorite-artist")
	{
		myFavoriteArtistServiceInterface = mfas
		myFavoriteArtist.GET("/top-tracks", GetTopTracks)
		myFavoriteArtist.GET("/artist-info", GetArtistInfo)
		myFavoriteArtist.GET("/image-info", GetImageInfo)
	}
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// GetTopTracks godoc
// @Summary      Show top tracks
// @Description  Get top tracks based on specified parameters
// @Tags         GetTopTracks
// @Accept       json
// @Produce      json
// @Param        method     query   string   true   "The method parameter"
// @Param        country    query   string   true   "The country parameter"
// @Param        api_key    query   string   true   "The API key parameter"
// @Param        format     query   string   true   "The format parameter"
// @Success      200  {Return raw Object of TopTracks }  myFavoriteArtistServiceInterface MyFavoriteArtistService
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /top-tracks  [get]
func GetTopTracks(c *gin.Context) {

	// Extracting parameters from the query
	method := c.Query("method")
	countryName := c.Query("country")
	apiKey := c.Query("api_key")
	format := c.Query("format")
	lastFMAPIURL := fmt.Sprintf("%s"+"?method=%s"+"&country=%s"+"&api_key=%s"+"&format=%s", fmBaseURL, method, countryName, apiKey, format)
	topTracks, err := myFavoriteArtistServiceInterface.TopTracks(c, lastFMAPIURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"top_tracks": topTracks})

}

// GetArtistInfo godoc
// @Summary      Show Artist info
// @Description  Get top tracks based on specified parameters
// @Tags         GetArtistInfo
// @Accept       json
// @Produce      json
// @Param        method     query   string   true   "The method parameter"
// @Param        artist    query   string   true   "The country parameter"
// @Param        api_key    query   string   true   "The API key parameter"
// @Param        format     query   string   true   "The format parameter"
// @Success      200  {Return raw Object of TopTracks }  myFavoriteArtistServiceInterface MyFavoriteArtistService
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /artist-info  [get]
func GetArtistInfo(c *gin.Context) {
	method := c.Query("method")
	artistName := c.Query("artist")
	apiKey := c.Query("api_key")
	format := c.Query("format")
	lastFMAPIURL := fmt.Sprintf("%s"+"?method=%s"+"&artist=%s"+"&api_key=%s"+"&format=%s", fmBaseURL, method, artistName, apiKey, format)
	artistInfo, err := myFavoriteArtistServiceInterface.ArtistInfo(c, lastFMAPIURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"artistInfo": artistInfo})
}

// GetImageInfo godoc
// @Summary      Show image info
// @Description  Get top tracks based on specified parameters
// @Tags         GetImageInfo
// @Accept       json
// @Produce      json
// @Param        method     query   string   true   "The method parameter"
// @Param        artist    query   string   true   "The country parameter"
// @Param        api_key    query   string   true   "The API key parameter"
// @Param        format     query   string   true   "The format parameter"
// @Success      200  {Return raw Object of TopTracks }  myFavoriteArtistServiceInterface MyFavoriteArtistService
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /image-info  [get]
func GetImageInfo(c *gin.Context) {

	method := c.Query("method")
	artistName := c.Query("artist")
	apiKey := c.Query("api_key")
	format := c.Query("format")
	lastFMAPIURL := fmt.Sprintf("%s"+"?method=%s"+"&artist=%s"+"&api_key=%s"+"&format=%s", fmBaseURL, method, artistName, apiKey, format)
	imageInfo, err := myFavoriteArtistServiceInterface.ImageInfo(c, lastFMAPIURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"imageInfo": imageInfo})

}
