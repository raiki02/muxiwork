package db

import "fmt"

func InsertUser(id, name, pid, label string) {
	sql := `insert into getname(id ,name ,pid ,label) values( ? ,? ,? ,? )`
	res, err := DB.Exec(sql, id, name, pid, label)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	newid, err := res.LastInsertId()
	if err != nil {
		fmt.Println("get lastinsert ID failed, ", err)
		return

	}
	fmt.Println("insert succ:", newid)
}
