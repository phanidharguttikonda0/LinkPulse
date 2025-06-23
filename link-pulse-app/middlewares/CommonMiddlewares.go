package middlewares

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/services"
	"log"
	"net/http"
)

func IsPremiumCheck(db *sql.DB, number string) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(int)
		value, err := services.CheckPremium(db, userId, number)
		if err != nil {
			log.Println("error was ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		if value {
			c.Set("premium", true)
		} else {
			c.Set("premium", false)
		}
		c.Next()
	}
}
