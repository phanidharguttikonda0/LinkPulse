package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/models"
	"log"
	"net/http"
)

const CustomNameRegex = `^[a-zA-Z0-9]{3,25}$`

func CustomNameValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var customName models.CustomName
		err := c.ShouldBind(&customName)
		if err != nil {
			log.Printf("ShouldBind: %v\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		value := CustomNameValidation(customName.Name)
		if value {
			c.Set("CustomName", customName.Name)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "CustomName validation failed"})
		}
	}
}

func CustomNameValidationMiddlewareGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		CustomName := c.Param("name")
		value := CustomNameValidation(CustomName)
		if value {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "CustomName validation failed"})
		}
	}
}

func CustomNameValidation(customName string) bool {

	log.Println(customName, " Regex Checking was going On")
	if models.IsValid(CustomNameRegex, customName) {
		return true
	} else {
		return false
	}

}
