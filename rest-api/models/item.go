package models

import "time"

type OrderItem struct {
	//gorm.Model
	ID          uint      `gorm:"primaryKey" json:"-"`
	OrderID     uint      `json:"-"`
	ItemCode    string    `json:"itemCode"`
	Description string    `gorm:"not null;type:varchar(150)"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
