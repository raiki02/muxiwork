package model

import "gorm.io/gorm"

type Book struct {
	Author string `gorm:"column:author; not null" json:"author"`
	Name   string `gorm:"column:name; not null" json:"name"`
	Model  *gorm.Model
	IsLent bool `gorm:"column:is_lent; default:false; enum:'true;false'" json:"is_lent"`
}
