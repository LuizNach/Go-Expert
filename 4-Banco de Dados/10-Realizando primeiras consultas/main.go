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

	db.Create(&Product{ // Como o campo ID é primary key ele vai auto incrementar se vc não passar para ele.
		Name:  "Playstation",
		Price: 59.99,
	})

	// products := []Product{
	// 	{
	// 		Name:  "XBox",
	// 		Price: 499.99,
	// 	},
	// 	{
	// 		Name:  "Nintendo Switch",
	// 		Price: 199.99,
	// 	},
	// }
	// db.Create(products) // Insert em batch

	var product Product
	db.First(&product, 2)
	fmt.Println("O produto encontrado é:")
	fmt.Println(product)

	// db.First(&product, "name = ?", "Playstation")
	// // Como já foi feito um First se voce colocar outro First ele acumula os selectors ou seja
	// // vai procurar agora por SELECT * FROM `products` WHERE name = 'Playstation' AND `products`.`id` = 2 ORDER BY `products`.`id` LIMIT 1
	// fmt.Println("O produto encontrado é:")
	// fmt.Println(product)

	// db.First(&product, 0) // buscando por valor que não existe no banco
	// // Como já foi feito um First se voce colocar outro First ele acumula os selectors ou seja
	// // vai procurar agora por SELECT * FROM `products` WHERE `products`.`id` = 0 AND `products`.`id` = 2 ORDER BY `products`.`id` LIMIT 1
	// fmt.Println("O produto encontrado é:")
	// fmt.Println(product)

	fmt.Println()

	// select all
	var products []Product
	db.Find(&products) // fazendo o select all
	for _, value := range products {
		//fmt.Println(key)
		fmt.Println(value)
	}

}
