package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Age   int    `gorm:"not null"`
	Grade string `gorm:"size:50;not null"`
}

func create(db *gorm.DB) {
	student := Student{Name: "张三", Age: 20, Grade: "三年级"}
	result := db.Create(&student)

	if result.Error != nil {
		fmt.Printf("Error creating student: %v\n", result.Error)
		return
	}

	fmt.Printf("result:%v\n", student.ID)
}

func read(db *gorm.DB) {
	var students []Student
	result := db.Where("Age >= ?", "18").Find(&students) // 查找ID为1的学生

	if result.Error != nil {
		fmt.Printf("Error reading student: %v\n", result.Error)
		return
	}
	for _, student := range students {
		fmt.Printf("Student found: ID=%d, Name=%s, Age=%d, Grade=%s\n", student.ID, student.Name, student.Age, student.Grade)
	}
}

func update(db *gorm.DB) {
	var student Student

	student.Grade = "四年级"
	result := db.Model(&student).Where("Name = ?", "张三").
		Updates(student)

	if result.Error != nil {
		fmt.Printf("Error updating student: %v\n", result.Error)
		return
	}

	fmt.Printf("Updated student: ID=%d, Name=%s, Age=%d, Grade=%s\n", student.ID, student.Name, student.Age, student.Grade)
}

func delete(db *gorm.DB) {
	var students []Student
	result := db.Where("Age < ?", "15").Delete(&students)

	if result.Error != nil {
		fmt.Printf("Error deleting student: %v\n", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Println("No student found to delete.")
	} else {
		fmt.Println("RowsAffected: ", result.RowsAffected)
	}
}

func main() {
	user := "rwuser"
	pass := "Rwpass@123"
	dsn := user + ":" + pass + "@tcp(47.111.78.104:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	create(db)
	read(db)
	update(db)
	delete(db)
}
