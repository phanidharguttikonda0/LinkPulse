package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/db"
	_ "github.com/phanidharguttikonda0/LinkPulse/db"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	_ "github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
)

func main() {
	r := gin.Default()

	log.Println("<UNK> Connected to RDS successfully!")
	connection, jwtSecret := db.DatabaseConnections()
	// log.Println(connection)

	// let's establish the connection
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

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}
