package main

import (
	"github.com/gin-gonic/gin"
	"go-gorm/database"
	"log"
	"rest-api/database"
	"rest-api/handler"
)

func main() {
	database.StartDB()

	// gin
	router := gin.Default()

	//routes
	//router.POST("/orders", handler.createOrder(db))
	router.POST("/orders", handler.CreateOrder(database))
	router.GET("/orders", handler.GetOrders(database))
	router.PUT("/orders/:orderId", handler.UpdateOrder(database))
	router.DELETE("/orders/:orderId", handler.DeleteOrder(database))

	// start service
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
