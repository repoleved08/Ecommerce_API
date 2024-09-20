package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"unique"`
	Email      string `json:"password"`
	Password   string `json:"-"`
	Role       string `json:"role"`
}
