package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SignIn(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func SignUp(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {}
}
