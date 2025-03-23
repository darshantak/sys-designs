package kvstore

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var kvStore = make(map[string]string)
var kvMu sync.Mutex

type setKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func setKey(c *gin.Context) {
	kvMu.Lock()
	defer kvMu.Unlock()
	var req setKv
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "request body"})
		return
	}

	kvStore[req.Key] = req.Value
	c.JSON(http.StatusOK, gin.H{"message": "key stored successfully"})
}

func getKey(c *gin.Context) {
	kvMu.Lock()
	defer kvMu.Unlock()
	reqKey := c.Param("key")
	if _, exists := kvStore[reqKey]; !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "key not found"})
	}

	c.JSON(http.StatusOK, gin.H{reqKey: kvStore[reqKey]})
}

func deleteKey(c *gin.Context) {
	kvMu.Lock()
	defer kvMu.Unlock()
	reqKey := c.Param("key")
	if _, exists := kvStore[reqKey]; !exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "key not found for deletion"})
		return
	}
	delete(kvStore, reqKey)
	c.JSON(http.StatusOK, gin.H{"message": "key successfully deleted"})
}

func KvMain(r *gin.Engine) {
	r.POST("/set", setKey)
	r.GET("/get/:key", getKey)
	r.DELETE("/delete/:key", deleteKey)

}
