package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}) // Migração para criar um tabela a partir dessa struct

	db.Create(&Product{ // Como o campo ID é primary key ele vai auto incrementar se vc não passar para ele.
		Name:  "PS2",
		Price: 99.99,
	})

	products := []Product{
		{
			Name:  "XBox",
			Price: 499.99,
		},
		{
			Name:  "Nintendo Switch",
			Price: 199.99,
		},
	}
	db.Create(products) // Insert em batch

}
