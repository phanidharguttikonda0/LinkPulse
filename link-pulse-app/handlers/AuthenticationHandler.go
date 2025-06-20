package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
	"github.com/phanidharguttikonda0/LinkPulse/models"
	"github.com/phanidharguttikonda0/LinkPulse/services"
	"log"
)

func SignIn(db *sql.DB, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Println("Sign In Handler, Going to Call Validations")
		var data models.User
		if err := c.ShouldBind(&data); err != nil {
			log.Printf("ShouldBind: %v\n", err)
			c.JSON(400, gin.H{"error": err})
			return
		}

		_, err := data.SignInValidation()

		if err != nil {
			log.Printf("Validation error : %v\n", err)
			c.JSON(400, gin.H{"error": err})
			return
		}
		log.Println("Validation was Completed Successfully")
		log.Println("Going to Call the Sign In Service")
		value, id := services.CheckUser(db, &data)
		if !value {
			c.JSON(400, gin.H{"error": "Invalid Credentials"})
			return
		}
		log.Println("Sign In Service Successfully, Going to Authorization header")
		authorizationHeader, err := middlewares.CreateAuthorizationHeader(jwtSecret, id, data.Username)
		if err != nil {
			log.Printf("Authorization error : %v\n", err)
			c.JSON(400, gin.H{"error": err})
			return
		}
		c.Header("Authorization", authorizationHeader)
		log.Println("Authorization Header was created Successfully")
		c.JSON(200, gin.H{"message": "Authorized"})
	}
}

func SignUp(db *sql.DB, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Sign Up Handler, Going to Call Validations")
		var data models.NewUser
		if err := c.ShouldBind(&data); err != nil {
			log.Printf("ShouldBind: %v\n", err)
			c.JSON(400, gin.H{"error": err})
			return
		}
		log.Println("Validation was Completed Successfully")
		log.Println("Going to store the new user")
		value, id := services.NewUser(db, &data)

		if !value {
			c.JSON(400, gin.H{"error": "User already exists with some Credentials"})
			return
		}

		log.Println("Stored User Successfully")
		log.Println("Going to get Authorization Header")

		authorizationHeader, err := middlewares.CreateAuthorizationHeader(jwtSecret, id, data.User.Username)
		if err != nil {
			log.Printf("Authorization error : %v\n", err)
			c.JSON(400, gin.H{"error": err})
			return
		}

		c.Header("Authorization", authorizationHeader)
		log.Println("Authorization Header was created Successfully")
		c.JSON(200, gin.H{"message": "Authorized"})
	}
}
