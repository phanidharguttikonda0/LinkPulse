package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/services"
	"log"
)

func IsPremium(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		userId := c.MustGet("userId").(int)
		number := c.Param("num")
		value, err := services.CheckPremium(db, userId, number)
		if err != nil {
			log.Println("CheckPremium Error:", err.Error())
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		if value == false {
			c.JSON(200, gin.H{
				"isPremium": false,
			})
		} else {
			c.JSON(200, gin.H{
				"isPremium": true,
			})
		}
	}
}
