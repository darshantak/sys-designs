package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ShortenRequest struct {
	Url string `json:"url"`
}

type ShortenResponse struct {
	ShortUrl string `json:"newUrl"`
}

var (
	urlStorage = make(map[string]string)
	mutex      = sync.Mutex{}
)

func shortenURL(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	shortCode := uuid.New().String()[:6]
	urlStorage[shortCode] = req.Url
	c.JSON(http.StatusOK, ShortenResponse{ShortUrl: fmt.Sprintf("http://localhost:8181/%s", shortCode)})
}

func redirectUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, exists := urlStorage[shortCode]

	if exists {
		c.Redirect(http.StatusFound, url)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
	}
}

func deleteUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if _, exists := urlStorage[shortCode]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "This url doesn't exists"})
		return
	} else {
		delete(urlStorage, shortCode)
		c.JSON(http.StatusNotFound, gin.H{"message": "URL deleted"})
	}
}
func listUrl(c *gin.Context) {
	c.JSON(http.StatusOK, urlStorage)
}

func UrlMain() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/shorten", shortenURL)
	r.GET("/:shortCode", redirectUrl)
	r.DELETE("/:shortCode", deleteUrl)
	r.GET("/list", listUrl)
	r.Run(":8181")
}
