package kvstore

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to set up the router
func setupRouter() *gin.Engine {
	r := gin.Default()
	KvMain(r)
	return r
}

// Test setting a key
func TestSetKey(t *testing.T) {
	r := setupRouter()

	body := map[string]string{"key": "testKey", "value": "testValue"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Test getting a key
func TestGetKey(t *testing.T) {
	r := setupRouter()

	// First, set a key
	body := map[string]string{"key": "testKey", "value": "testValue"}
	jsonBody, _ := json.Marshal(body)

	reqSet, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonBody))
	reqSet.Header.Set("Content-Type", "application/json")

	wSet := httptest.NewRecorder()
	r.ServeHTTP(wSet, reqSet)

	// Now, get the key
	reqGet, _ := http.NewRequest("GET", "/get/testKey", nil)
	wGet := httptest.NewRecorder()
	r.ServeHTTP(wGet, reqGet)

	assert.Equal(t, http.StatusOK, wGet.Code)
}

// Test deleting a key
func TestDeleteKey(t *testing.T) {
	r := setupRouter()

	// First, set a key
	body := map[string]string{"key": "testKey", "value": "testValue"}
	jsonBody, _ := json.Marshal(body)

	reqSet, _ := http.NewRequest("POST", "/set", bytes.NewBuffer(jsonBody))
	reqSet.Header.Set("Content-Type", "application/json")

	wSet := httptest.NewRecorder()
	r.ServeHTTP(wSet, reqSet)

	// Now, delete the key
	reqDelete, _ := http.NewRequest("DELETE", "/delete/testKey", nil)
	wDelete := httptest.NewRecorder()
	r.ServeHTTP(wDelete, reqDelete)

	assert.Equal(t, http.StatusOK, wDelete.Code)
}

// Test syncing all keys
