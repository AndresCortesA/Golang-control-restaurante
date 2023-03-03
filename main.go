package main

import (
	_ "golang-RestauranteControl/middleware"
	_ "golang-RestauranteControl/routes"
	_ "golang-RestauranteControl]/database"
	"os"

	_ "github.com/gin-gonic/gin"
	_ "go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.new()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)
	router.Run(":" + port)

}
