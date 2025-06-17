package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
)

func AuthenticationRoutes(r *gin.Engine, db *sql.DB) {
	authenticationRoute := r.Group("/authentication")
	authenticationRoute.POST("/sign-up", handlers.SignUp(db))
	authenticationRoute.POST("/sign-in", handlers.SignIn(db))
}
