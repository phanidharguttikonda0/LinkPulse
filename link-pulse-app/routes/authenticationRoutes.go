package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
)

func AuthenticationRoutes(r *gin.Engine, db *sql.DB, jwtSecret string) {
	authenticationRoute := r.Group("/authentication")
	authenticationRoute.POST("/sign-up", handlers.SignUp(db, jwtSecret))
	authenticationRoute.POST("/sign-in", handlers.SignIn(db, jwtSecret))
}
