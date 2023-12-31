package routes

import (
	"e-commerce-cart/controllers"
	"github.com/gin-gonic/gin"
)

// Handling routes
func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/login", controllers.SignUp())
	incomingRoutes.POST("/users/signup", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
