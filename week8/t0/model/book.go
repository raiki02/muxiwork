package model

type BookInfo struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Author   string `json:"author" validate:"required"`
	Category string `json:"category" validate:"required"`
	IsLent   bool   `json:"is_lent" validate:"required"`
}
