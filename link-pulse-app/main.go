package main

import (
	""
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	_ "github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	routes.AuthenticationRoutes(r)

	r.GET("/", func(c *gin.Context) {
		log.Println("Called the base resource")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Gin!",
		})
	})

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}
