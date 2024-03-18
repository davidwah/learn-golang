package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	_models "rest-api/models"
	"strconv"
)

type OrderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

// fungsi CreateOrder
func (c *OrderHandler) CreateOrder(ctx *gin.Context) {
	var order _models.Order
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save order and order items
	if err := c.db.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

// fungsi GetOrders
func (c *OrderHandler) GetOrders(ctx *gin.Context) {
	var orders []_models.Order
	if err := c.db.Preload("Items").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// fungsi UpdateOrder
func (c *OrderHandler) UpdateOrder(ctx *gin.Context) {
	id := ctx.Param("orderId")
	var order _models.Order
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update order and items
	if err := c.db.Model(&_models.Order{ID: toUint(id)}).Updates(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch updated order with items
	if err := c.db.Preload("Items").First(&order, "id = ?", id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// fungsi DeleteOrder
func (c *OrderHandler) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("orderId")
	if err := c.db.Delete(&_models.Order{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Success delete"})
}

func toUint(str string) uint {
	id, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return uint(id)
}
