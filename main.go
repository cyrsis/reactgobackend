package main

import (
	"net/http"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files in the Views Folder
	router.Use(static.Serve("/", static.LocalFile("./Views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}



	// Start and run the server
	router.Run(":3000")
}
