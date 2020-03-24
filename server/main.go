package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/gmlwo530/maru-web-app-project/controllers"
	myDB "github.com/gmlwo530/maru-web-app-project/db"
)

func main() {
	mode := os.Getenv("GIN_MODE")
	g := gin.Default()

	g.Use(static.Serve("/", static.LocalFile("./web", true)))
	

	db := myDB.SetupModels()

	// Provide db variable to controllers
	// Create a middleware that can provide the database instance to every single controller since
	// they live in another file that can't access the database instance directly
	g.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	api := g.Group("/api")
	api.GET("/ping", controllers.GetPing)

	if mode == "release" {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("$PORT must be set")
		}
		log.Fatal(http.ListenAndServe(":" + port, g))
	} else {
		port := ":8080"
		log.Fatal(http.ListenAndServe(port, g))
	}
}