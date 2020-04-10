package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gmlwo530/maru-web-app-project/db/models"
	"github.com/stretchr/testify/assert"
)

var newID string
var router *gin.Engine = SetGin()

func TestCreatePost(t *testing.T) {
	newTitle := "new title"
	newContent := "new content"

	s := fmt.Sprintf(`{"title":"%s","content":"%s"}`, newTitle, newContent)

	jsonStr := []byte(s)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/posts", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resultPost models.Post
	if err := json.NewDecoder(w.Body).Decode(&resultPost); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, newTitle, resultPost.Title)
	assert.Equal(t, newContent, resultPost.Content)
	newID = resultPost.IdtoA()
}

func TestGetAllPost(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/posts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resultPosts []models.Post
	if err := json.NewDecoder(w.Body).Decode(&resultPosts); err != nil {
		log.Fatal(err)
	}
}

func TestGetPost(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/posts/"+newID, nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var resultPost models.Post
	if err := json.NewDecoder(w.Body).Decode(&resultPost); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newID, resultPost.IdtoA())
}

func TestUpdatePost(t *testing.T) {
	updatedTitle := "update title"
	updatedContent := "update content"

	s := fmt.Sprintf(`{"title":"%s","content":"%s"}`, updatedTitle, updatedContent)

	var jsonStr = []byte(s)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/posts/"+newID, bytes.NewBuffer(jsonStr))

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var resultPost models.Post
	if err := json.NewDecoder(w.Body).Decode(&resultPost); err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, updatedTitle, resultPost.Title)
	assert.Equal(t, updatedContent, resultPost.Content)
}

func TestDeletePost(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/posts/"+newID, nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/posts/"+newID, nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
