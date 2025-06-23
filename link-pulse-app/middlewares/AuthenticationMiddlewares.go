package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/phanidharguttikonda0/LinkPulse/models"
	"log"
	"net/http"
	"time"
)

func AuthorizationCheckMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		AuthorizationHeader := c.GetHeader("Authorization")
		claims, err := AuthorizationCheck(secret, AuthorizationHeader)
		if err != nil {
			log.Printf("AuthorizationCheck: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			return
		}
		username, ok := claims["username"].(string)
		if !ok {
			log.Println("Unable to get username from the claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			return
		}

		userId, ok := claims["userId"].(int)

		if !ok {
			log.Println("Unable to get userId from the claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
			return
		}
		c.Set("username", username) // username := c.MustGet("username").(string) to get the value
		c.Set("userId", userId)
		c.Next()
	}
}

func AuthorizationCheck(secret string, authorizationHeader string) (jwt.MapClaims, error) {
	jwtSecret := []byte(secret)
	token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
		// validating signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func CreateAuthorizationHeader(secret string, userId int, username string) (string, error) {
	jwtSecret := []byte(secret)
	claims := jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func SignInValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.User
		if err := c.ShouldBind(&data); err != nil {
			log.Printf("ShouldBind: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		_, err := data.SignInValidation()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		c.Set("data", data)
		log.Println("Validation was successful")
		c.Next()
	}
}

func SignUpValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data models.NewUser
		if err := c.ShouldBind(&data); err != nil {
			log.Printf("ShouldBind: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		_, err := data.SignUpValidation()

		if err != nil {
			log.Printf("SignUpValidation error: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}

		c.Set("data", data)
		log.Println("Validation was successful")
		c.Next()
	}
}
