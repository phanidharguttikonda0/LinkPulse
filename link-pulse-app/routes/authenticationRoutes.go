package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phanidharguttikonda0/LinkPulse/handlers"
)

func AuthenticationRoutes(r *gin.Engine) {
	authenticationRoute := r.Group("/authentication")
	authenticationRoute.POST("/sign-up", handlers.SignUp)
	authenticationRoute.POST("/sign-in", handlers.SignIn)
}
