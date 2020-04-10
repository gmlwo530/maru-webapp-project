package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gmlwo530/maru-web-app-project/db/models"
	"github.com/jinzhu/gorm"
)

func createPing(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	ping := models.Ping{CreatedAt: time.Now()}
	db.Create(&ping)
}

// GetPing is get ping
func GetPing(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	defer createPing(c)

	var ping models.Ping

	message := "first time!"

	if !db.Order("created_at").Last(&ping).RecordNotFound() {
		message = fmt.Sprintf("%v ago", time.Now().Sub(ping.CreatedAt).String())
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
