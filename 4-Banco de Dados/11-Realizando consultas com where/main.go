package main

import (
	"fmt"

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

	// db.Create(&Product{ // Como o campo ID é primary key ele vai auto incrementar se vc não passar para ele.
	// 	Name:  "MegaDrive",
	// 	Price: 49.99,
	// })

	fmt.Println()

	// select all
	var products []Product
	db.Find(&products) // fazendo o select all
	for _, value := range products {
		//fmt.Println(key)
		fmt.Println(value)
	}

	fmt.Println()

	// select como limit 2
	fmt.Println("Fazendo o select do gorm com limite:2")
	db.Limit(2).Find(&products) // fazendo o select all
	for _, value := range products {
		//fmt.Println(key)
		fmt.Println(value)
	}

	fmt.Println()

	fmt.Println("Fazendo o select do gorm com limite:2 e offset:2")
	// select com paginação ou seja é limitado a dois por página mas queremos o offset(numero da pagina 2)
	// entao resgatamos os indices 3 e 4 na ordem do resultado do select all
	db.Limit(2).Offset(2).Find(&products) // fazendo o select all
	for _, value := range products {
		//fmt.Println(key)
		fmt.Println(value)
	}

	fmt.Println()

	// where
	fmt.Println("Fazendo o select do gorm com where price > 50")
	db.Where("price > ?", 50).Find(&products)
	for _, value := range products {
		fmt.Println(value)
	}

	fmt.Println()
	// like
	fmt.Println("Fazendo o select do gorm com like")
	db.Where("name like ?", "%station%").Find(&products) // aqui buscamos todos os items que possuem o name que contenha a substring "station"
	for _, value := range products {
		fmt.Println(value)
	}
}
