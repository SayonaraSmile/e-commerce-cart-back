package main

import (
	"e-commerce-cart/controllers"
	"e-commerce-cart/database"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),
		database.UserData(database.Client("Users")))
}
