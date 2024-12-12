package service

import (
	"log"
	"t0/internal/model"
	"time"
)

func Register(username, password string) {
	now := time.Now()
	user := model.User{
		UserName: username,
		Password: password,
	}
	user.Model.CreatedAt = now

	res := DB.Create(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Println("User registered, Id: ", user.Model.ID)

}

func UpdatePassword(username, old_password, new_password string) {
	if old_password == new_password {
		log.Fatal("Old password and new password are the same")
	}

	user, err := FindUser(username)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = new_password
	res := DB.Save(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Println("Password updated")

}

func DeleteUser(username, password string) {
	user, err := FindUser(username)
	if err != nil {
		log.Fatal(err)
	}
	res := DB.Delete(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Println("User deleted")
}

func FindUser(username string) (user model.User, err error) {
	res := DB.Where("username = ?", username).First(&user)
	if res.Error != nil {
		log.Fatal(res.Error)
		return model.User{}, res.Error
	}
	log.Println("User found, Id: ", user.Model.ID)
	return user, nil
}
