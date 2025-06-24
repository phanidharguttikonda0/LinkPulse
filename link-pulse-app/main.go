package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/db"
	_ "github.com/phanidharguttikonda0/LinkPulse/db"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/routes"
	_ "github.com/phanidharguttikonda0/LinkPulse/routes"
	"github.com/phanidharguttikonda0/LinkPulse/services"
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
	routes.WebRoutes(r, connection, jwtSecret)
	routes.CommonRoutes(r, connection, jwtSecret)

	// number 1 represents whether the user has premium for the website insights
	r.GET("/:name", middlewares.AuthorizationCheckMiddleware(jwtSecret), middlewares.CustomNameValidationMiddlewareGet(), middlewares.IsPremiumCheck(connection, "1"), RedirectUrl(connection))
	// name is nothing but the shorten url

	// Start server
	r.Run(":8080") // default listens on 0.0.0.0:8080
}

func RedirectUrl(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		premium := c.MustGet("premium").(bool)
		url := c.Param("name")
		log.Println("got the Url", url)
		// here we are going to just increment the count that's it not more than that
		original, err := services.GetOriginalUrl(db, url)
		if err != nil {
			log.Printf("failed to retrieve originalUrl: %v because of the following ", err)
			c.Redirect(302, "/")
			return
		}

		c.Redirect(302, original)
		// after re-direction we will store remaining things such that latency will improve
		if premium {
			log.Println("It's a Premium Users Link")
			// here we are going get the Ip address and all the other stuff
			services.StoringPremiumUserInsights()
		}
		return
	}
}
