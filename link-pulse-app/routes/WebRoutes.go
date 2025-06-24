package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
)

func WebRoutes(r *gin.Engine, db *sql.DB, jwtSecret string) {
	webRoutes := r.Group("/website")
	webRoutes.GET("/url-shortner", middlewares.AuthorizationCheckMiddleware(jwtSecret), handlers.UrlShortner(db))
	webRoutes.POST("/url-shortner", middlewares.AuthorizationCheckMiddleware(jwtSecret),
		middlewares.CustomNameValidationMiddleware(), handlers.PostUrlShortner(db))
	webRoutes.GET("/custom-check/:name", middlewares.AuthorizationCheckMiddleware(jwtSecret),
		middlewares.CustomNameValidationMiddlewareGet(), handlers.CheckCustomName(db))
}
