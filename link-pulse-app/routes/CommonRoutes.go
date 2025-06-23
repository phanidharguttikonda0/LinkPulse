package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
)

// these are routes that are common for website and file sharing

func CommonRoutes(r *gin.Engine, db *sql.DB, secret string) {
	routes := r.Group("/common")
	routes.GET("/is-premium/:num", middlewares.AuthorizationCheckMiddleware(secret), handlers.IsPremium(db))
}
