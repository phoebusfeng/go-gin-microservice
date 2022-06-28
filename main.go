package main

import (
	"github.com/gin-gonic/gin"
	a "github.com/phoebusfeng/go-gin-microservice/handlers"
)

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", a.ShowIndexPage)
	router.GET("/article/view/:article_id", a.GetArticle)

	// Start serving the application
	router.Run(":9999")
}
