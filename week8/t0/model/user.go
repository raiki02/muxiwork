package model

import "gorm.io/gorm"

type User struct {
	UserName string `gorm:"column:username; unique" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Model    *gorm.Model
}
