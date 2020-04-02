package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gmlwo530/maru-web-app-project/db/models"
	"github.com/stretchr/testify/assert"
)

func TestPostRoute(t *testing.T) {
	router := SetGin()

	post := models.Post{
		Title:   "My Title",
		Content: "My Content",
	}

	jsonPost, _ := json.Marshal(post)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/posts", bytes.NewBuffer(jsonPost))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resultPost models.Post
	if err := json.NewDecoder(w.Body).Decode(&resultPost); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, post.Title, resultPost.Title)
}
