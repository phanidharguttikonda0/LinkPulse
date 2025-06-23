package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/db"
	_ "github.com/phanidharguttikonda0/LinkPulse/db"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	_ "github.com/phanidharguttikonda0/LinkPulse/routes"
	"log"
)

func main() {

	header, errr := middlewares.CreateAuthorizationHeader("phani", 1, "Phanidhar")
	if errr != nil {
		log.Fatal(errr)
	} else {
		fmt.Println("header: ", header)
	}
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
	routes.WebRoutes(r, connection, jwtSecret)
	routes.CommonRoutes(r, connection, jwtSecret)

	// number 1 represents whether the user has premium for the website insights
	r.GET("/:urlName", middlewares.AuthorizationCheckMiddleware(jwtSecret), middlewares.IsPremiumCheck(connection, "1"), RedirectUrl(connection))

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}

func RedirectUrl(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		premium := c.MustGet("premium").(bool)
		url := c.Param("urlName")
		log.Println("got the Url", url)
		// here we are going to just increment the count that's it not more than that

		if premium {
			// here we are going get the Ip address and all the other stuff
		}
		// here we will finally do the re-direction
	}
}
