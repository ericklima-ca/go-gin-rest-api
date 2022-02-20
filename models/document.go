package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Quantity uint   `gorm:"default:1" json:"quantity,omitempty"`
	Product  string `json:"product,omitempty"`
	OrderID  uint   `json:"order_id,omitempty"`
}
