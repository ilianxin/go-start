package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID      int     `gorm:"primaryKey"`
	Balance float64 `gorm:"not null"`
}

type Transaction struct {
	ID            int     `gorm:"primaryKey"`
	FromAccountId int     `gorm:"not null"`
	ToAccountId   int     `gorm:"not null"`
	Amount        float64 `gorm:"not null"`
}

func main() {
	user := "rwuser"
	pass := "Rwpass@123"
	dsn := user + ":" + pass + "@tcp(47.111.78.104:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 开启事务
	tx := db.Begin()
	if err := tx.Error; err != nil {
		panic("failed to begin transaction")
	}
	// 创建学生

	transaction := Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100.0}
	fromAccount := Account{ID: 1}
	db.First(&fromAccount, transaction.FromAccountId)
	if fromAccount.Balance < transaction.Amount {
		tx.Rollback()
		panic("insufficient balance")
	} else {
		fromAccount.Balance -= transaction.Amount
		toAccount := Account{ID: transaction.ToAccountId}
		db.First(&toAccount, transaction.ToAccountId)
		toAccount.Balance += transaction.Amount

		if err := tx.Save(&fromAccount).Error; err != nil {
			tx.Rollback()
			panic("failed to update from account balance")
		}
		if err := tx.Save(&toAccount).Error; err != nil {
			tx.Rollback()
			panic("failed to update to account balance")
		}

		if err := tx.Create(&transaction).Error; err != nil {
			tx.Rollback()
			panic("failed to create transaction")
		}
	}

	tx.Commit()
	fmt.Printf("Transaction successful: From Account ID %d to To Account ID %d, Amount: %.2f\n", transaction.FromAccountId, transaction.ToAccountId, transaction.Amount)

}
