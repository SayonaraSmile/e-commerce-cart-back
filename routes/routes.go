package routes

import (
	"github.com/gin-gonic/gin"
)

// Handling routes
func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/login")
	incomingRoutes.POST("users/signup")
}
