package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
