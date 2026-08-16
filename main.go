package main

import (
	"github.com/coderboy53/url-shortener/handlers"
	"github.com/coderboy53/url-shortener/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// initialize the gin Engine
	r := gin.Default()

	// set up routes with handlers
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		})
	})
	r.POST("/create-short-url", handlers.CreateShortUrl)
	r.GET("/:short-url", handlers.RedirectToInitialLink)

	// initialize redis store
	store.InitializeStore()

	// run the web server
	err := r.Run(":9808")
	if err != nil {
		logrus.Error(err)
	}
}
