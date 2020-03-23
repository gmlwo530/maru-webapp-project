package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	// "github.com/gmlwo530/stack-api-server/controllers"
	// "github.com/gmlwo530/stack-api-server/models"
)

func main() {
	mode := os.Getenv("GIN_MODE")
	g := gin.Default()

	g.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := g.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// db := models.SetupModels()

	// Provide db variable to controllers
	// Create a middleware that can provide the database instance to every single controller since
	// they live in another file that can't access the database instance directly
	// g.Use(func(c *gin.Context) {
	// 	c.Set("db", db)
	// 	c.Next()
	// })

	// g.GET("/folders", controllers.FindFolders)
	// g.POST("/folders", controllers.CreateFolder)
	// g.GET("/folders/:id", controllers.FindFolder)
	// g.PATCH("/folders/:id", controllers.UpdateFolder)
	// g.DELETE("/folders/:id", controllers.DeleteFolder)

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