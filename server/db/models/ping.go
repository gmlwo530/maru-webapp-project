package models

import (
	"time"
	_ "github.com/jinzhu/gorm" // import gorm
)

// Ping is ping
type Ping struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
}
