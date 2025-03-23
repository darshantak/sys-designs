package main

import (
	kvstore "webapps/kvStore"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	kvstore.KvMain(r)
	r.Run(":8181")
}

// POST /shorten {"url":"https://google.com"}
// GET /shortCode
