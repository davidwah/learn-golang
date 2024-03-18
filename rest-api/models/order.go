package models

import (
	"time"
)

type Order struct {
	//gorm.Model
	ID           uint      `gorm:"primaryKey" json:"id"`
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}
