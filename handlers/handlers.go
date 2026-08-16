package handlers

import (
	"net/http"

	"github.com/coderboy53/url-shortener/shortener"
	"github.com/coderboy53/url-shortener/store"
	"github.com/gin-gonic/gin"
)

func CreateShortUrl(c *gin.Context) {
	initialUrl := c.PostForm("initialUrl")
	userId := c.PostForm("user_id")
	if initialUrl == "" || userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User ID or Input URL is missing",
		})
		return
	}
	shortLink := shortener.GenerateShortLink(initialUrl, userId)
	c.JSON(
		http.StatusCreated, 
		gin.H {
			"short-link": "http://localhost9808"+shortLink,
		})
	store.StoreMapping(shortLink, initialUrl, userId)
}

func RedirectToInitialLink(c *gin.Context) {

}