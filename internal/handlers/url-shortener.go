package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"url-shortener-service/internal/database"
	"url-shortener-service/internal/models"
	"url-shortener-service/internal/repository"
	"url-shortener-service/pkg/utils"
)

// UrlShortener handles the URL shortening request.
func UrlShortener(c *gin.Context) {
	var req models.UrlShortenerReq

	// Bind the JSON request body to the UrlShortenerReq struct
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid request body")
		return
	}

	// Generate the short URL
	shortUrl, err := GenerateShortUrl(req)
	if err != nil {
		log.Printf("Something went wrong: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Store the short URL in the database
	storedUrl, err := StoreShortUrl(req.Url, shortUrl)
	if err != nil {
		log.Printf("Something went wrong: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Respond with the stored short URL
	c.JSON(http.StatusOK, gin.H{"shortener-url": storedUrl})
}

// GenerateShortUrl generates a short URL based on the provided request.
func GenerateShortUrl(req models.UrlShortenerReq) (string, error) {
	shortUrl := "localhost:8080/"

	lengthOfShortUrl := req.Length

	// Generate a custom path for the short URL
	customPath, err := utils.GenerateCustomRelativePath(lengthOfShortUrl)
	if err != nil {
		return "", err
	}

	shortUrl += customPath
	return shortUrl, nil
}

// StoreShortUrl stores the original URL and short URL in the database.
func StoreShortUrl(originalUrl, shortUrl string) (string, error) {
	urlStorage := database.UrlStorage{
		OriginalUrl:  originalUrl,
		ShortenerUrl: shortUrl,
	}

	urlRepository := repository.NewUrlsRepository(database.GetDbInstance())

	// Store the URL in the database
	storedUrl, err := urlRepository.Store(urlStorage)
	if err != nil {
		return "", err
	}

	return storedUrl.ShortenerUrl, nil
}

// RedirectUrl handles the redirection to the original URL based on the custom path.
func RedirectUrl(c *gin.Context) {
	customPath := c.Param("customPath")

	// Construct the short URL
	shortUrl := ConstructShortUrl(customPath)
	if shortUrl == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!!!"})
		return
	}

	// Find the original URL based on the short URL
	originalUrl, err := FindOriginalUrl(shortUrl)
	if err != nil {
		log.Printf("Something went wrong: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!!!"})
		return
	}

	// Redirect to the original URL
	c.Redirect(http.StatusFound, originalUrl)
}

// ConstructShortUrl constructs the short URL based on the custom path.
func ConstructShortUrl(customPath string) string {
	shortUrl := "localhost:8080/"
	shortUrl += customPath
	return shortUrl
}

// FindOriginalUrl finds the original URL based on the short URL.
func FindOriginalUrl(shortUrl string) (string, error) {
	urlRepository := repository.NewUrlsRepository(database.GetDbInstance())
	url, err := urlRepository.FindBy("shortener_url", shortUrl)
	if err != nil {
		return "", err
	}
	return url.OriginalUrl, nil
}
