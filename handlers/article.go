package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	a "github.com/phoebusfeng/go-gin-microservice/models"
)

func ShowIndexPage(c *gin.Context) {
	articles := a.GetAllArticles()

	// Call the HTML method of the Context to render a template
	Render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")

}

func GetArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := a.GetArticleByID(articleID); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
