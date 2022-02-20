package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Documents []Document `json:"documents,omitempty"`
}
