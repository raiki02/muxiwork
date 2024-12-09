package model

type UserInfo struct {
	ID       string `json:"id" db:"id" validate:"required"`
	UserName string `json:"username" db:"username" validate:"required"`
	PassWord string `json:"password" db:"password" validate:"required"`
}

type UpdateUserInfo struct {
	ID           string `json:"id" db:"id" validate:"required"`
	UserName     string `json:"username" db:"username" validate:"required"`
	Old_PassWord string `json:"old_password" db:"password" validate:"required"`
	New_PassWord string `json:"new_password" db:"password" validate:"required"`
}
