package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     string `db:"salary"`
}

var db *sqlx.DB

func main() {
	user := "rwuser"
	pass := "Rwpass@123"

	var err error
	db, err = sqlx.Open("mysql", user+":"+pass+"@tcp(47.111.78.104:3306)/go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库异常:", err)
		return
	}

	var users []Employee
	db.Select(&users, "SELECT * FROM employees WHERE department = ?", "技术部")

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %s\n", user.ID, user.Name, user.Department, user.Salary)
	}

	var topSalary Employee

	err = db.Get(&topSalary, "SELECT * FROM employees order by salary desc Limit 1")

	if err != nil {
		fmt.Printf("Error retrieving top salary employee: %v\n", err)
	} else {
		fmt.Printf("Top Salary Employee: ID: %d, Name: %s, Department: %s, Salary: %s\n", topSalary.ID, topSalary.Name, topSalary.Department, topSalary.Salary)

	}

	defer db.Close()
}
