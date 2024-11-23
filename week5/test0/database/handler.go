package database

import "fmt"

func Insert(user_id, username, password string) {
	sql := `insert into users (user_id , username, password) values (?,  ? , ?)`
	res, err := DB.Exec(sql, user_id, username, password)
	if err != nil {
		fmt.Println("insert failed, err: ", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get last insert id failed, err: ", err)
		return
	}
	fmt.Println("insert success, the id is ", id)
}

func Update_password(username, old_password, new_password string) {
	sql := `update users set password = ? where username  = ? &&  password = ?`
	_, err = DB.Exec(sql, new_password, username, old_password)
	if err != nil {
		fmt.Println("update failed, err: ", err)
		return
	}
	fmt.Println("update successfully !")
}

func Query(user_id string) {
	sql := `select user_id , username , password from users where user_id = ?`
	_, err := DB.Exec(sql, user_id)
	if err != nil {
		fmt.Println("query failed, err: ", err)
		return
	}
	fmt.Println("show user successfully !")
}

func Delete_user(user_id string) {
	sql := `delete from users where user_id = ?`
	_, err := DB.Exec(sql, user_id)
	if err != nil {
		fmt.Println("delete failed, err: ", err)
		return
	}
	fmt.Println("delete user successfully !")
}
