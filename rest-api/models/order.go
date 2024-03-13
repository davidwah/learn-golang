package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	//gorm.Model
	ID           uint           `gorm:"primaryKey" json:"id"`
	OrderedAt    time.Time      `json:"orderedAt"`
	CustomerName string         `json:"customerName"`
	Items        []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
