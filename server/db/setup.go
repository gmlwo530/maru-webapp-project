package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // import postgres

	"fmt"
	"os"

	"github.com/gmlwo530/maru-web-app-project/db/models"
)

const (
	host    = "localhost"
	port    = 5432
	user    = "maru"
	dbname  = "maruwebapp"
	sslmode = "disable"
)

// SetupModels is set models
func SetupModels() *gorm.DB {
	mode := os.Getenv("GIN_MODE")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	var connStr string

	if mode == "release" {
		connStr = os.Getenv("DATABASE_URL")
	} else {
		connStr = "host=%s port=%d user=%s dbname=%s sslmode=%s password=%s"
		connStr = fmt.Sprintf(connStr, host, port, user, dbname, sslmode, dbPassword)
	}

	db, err := gorm.Open("postgres", connStr)

	if err != nil {
		fmt.Println("err : ", err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Ping{}, &models.Post{})

	return db
}
