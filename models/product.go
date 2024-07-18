package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"type:varchar(100);not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int     `json:"stock" gorm:"type:int(11);not null"`
	Category    string  `json:"category" gorm:"type:varchar(50);not null"`
}
