package user

type User struct {
	User_id  string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type Update_user struct {
	User_id      string `db:"usr_id" json:"usr_id"`
	Username     string `db:"username" json:"username"`
	Old_Password string `db:"old_password" json:"old_password"`
	New_Password string `db:"new_password" json:"new_password"`
}
