package dao

import (
	"log"
	"t0/model"
)

func AddBook(author, name string) {
	book := model.Book{
		Author: author,
		Name:   name,
	}
	res := DB.Create(&book)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Println("Book added Success, Id: ", book.Model.ID)

}

func BorrowBook(author, name string) {
	book := FindBook(author, name)
	if book.IsLent == true {
		log.Fatal("Book is already lent, try another book or just wait")
	} else {
		book.IsLent = true
		res := DB.Save(&book)
		if res.Error != nil {
			log.Fatal(res.Error)
		}
		log.Println("You borrowed the book ", book.Name)
	}
}

func DeleteBook(author, name string) {
	book := FindBook(author, name)
	res := DB.Delete(&book)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Printf("Book &s by &s deleted, Id: ", book.Name, book.Author, book.Model.ID)
}

func FindBook(author, name string) (book model.Book) {
	res := DB.Where("author = ? AND name = ? ", author, name).First(&book)
	if res.Error != nil {
		log.Fatal(res.Error)
	}
	log.Println("Book found, Id: ", book.Model.ID)
}
