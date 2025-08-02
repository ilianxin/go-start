package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	user := "rwuser"
	pass := "Rwpass@123"

	var db *sqlx.DB

	var err error
	db, err = sqlx.Open("mysql", user+":"+pass+"@tcp(47.111.78.104:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}

	var books []Book

	err = db.Select(&books, "SELECT * FROM books WHERE price > ?", 50.0)

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %.2f\n", book.ID, book.Title, book.Author, book.Price)
	}

	defer db.Close()
}
