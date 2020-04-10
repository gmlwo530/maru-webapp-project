package models

import (
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm" // import gorm
)

// Post is model
type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IdtoA is return id to string type
func (p Post) IdtoA() string {
	return strconv.FormatUint(uint64(p.ID), 10)
}
