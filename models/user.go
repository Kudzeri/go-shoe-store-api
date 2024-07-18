package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(50);unique;not null"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `json:"password" gorm:"type:text;not null"`
	FullName string `json:"full_name" gorm:"type:varchar(100)"`
	Age      int    `json:"age"`
}
