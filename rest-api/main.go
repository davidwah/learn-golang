package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"rest-api/database"
	"rest-api/handler"
	"rest-api/models"
)

func main() {
	db, err := database.StartDb()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models.Order{}, models.Item{})

	// gin service
	router := gin.Default()
	orderController := handler.NewOrderHandler(db)

	// router endpoint
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetOrders)
	router.PUT("/orders/:orderId", orderController.UpdateOrder)
	router.DELETE("/orders/:orderId", orderController.DeleteOrder)

	// start service
	erRoute := router.Run(":8080")
	if erRoute != nil {
		log.Fatal("Failed to start server:", erRoute)
	}
}
