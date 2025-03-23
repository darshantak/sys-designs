package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Limit struct {
	
}

func setLimit(c *gin.Context) {

}
func setResource(c *gin.Context) {

}
func RateMain() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/set-limit", setLimit)
	r.POST("/resource", setResource)

	r.Run(":8181")
}

// POST /api/resource → Accepts API requests only if the client has quota available.
// POST /admin/set-limit → Admin can set a custom rate limit for a specific API key.
