package main

import (
	"e-commerce-cart/controllers"
	"e-commerce-cart/database"
	"e-commerce-cart/middleware"
	"e-commerce-cart/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),
		database.UserData(database.Client, "Users")) // Creating new application

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoute(router)                // Handling routes
	router.Use(middleware.Authentication()) // Handling middleware

	// Handling routes
	router.GET("/addtocart"app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckoout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port)) // Running server
}
