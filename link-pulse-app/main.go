package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/db"
	_ "github.com/phanidharguttikonda0/LinkPulse/db"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	_ "github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.RateLimiterMiddleware()) // it will be called for each route before being executed
	log.Println("<UNK> Connected to RDS successfully!")
	connection, jwtSecret := db.DatabaseConnections()

	err := connection.Ping()
	if err != nil {
		log.Fatalf("failed to connect to RDS: %v", err)
	} else {
		log.Println("No Error Occurred")
	}

	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Fatalf("failed to close connection: %v", err)
		}
	}(connection)

	routes.AuthenticationRoutes(r, connection, jwtSecret) // for each route we are going to pass the database connection from here

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}
