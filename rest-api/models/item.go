package models

type Item struct {
	//gorm.Model
	ID          uint   `gorm:"primaryKey" json:"-"`
	OrderID     uint   `json:"-"`
	Code        string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
