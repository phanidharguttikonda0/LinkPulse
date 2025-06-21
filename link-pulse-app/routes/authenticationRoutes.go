package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
	"github.com/phanidharguttikonda0/LinkPulse/middlewares"
)

func AuthenticationRoutes(r *gin.Engine, db *sql.DB, jwtSecret string) {
	authenticationRoute := r.Group("/authentication")
	authenticationRoute.POST("/sign-up", middlewares.SignUpValidationMiddleware(), handlers.SignUp(db, jwtSecret))
	authenticationRoute.POST("/sign-in", middlewares.SignInValidationMiddleware(), handlers.SignIn(db, jwtSecret))
}
