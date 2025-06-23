package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func PostUrlShortner(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}

func UrlShortner(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {}
}

func CheckCustomName(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {}
}
