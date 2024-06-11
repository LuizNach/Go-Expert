package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

type Relationly struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &Relationly{}) // Adicionando Category para rodar nas migrações

	// Criando uma nova categoria
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// Criando um novo produto nessa categoria
	new_product := Product{Name: "Polystation", Price: 49.90, CategoryID: category.ID}
	db.Create(&new_product)

	var products []Product
	db.Find(&products)

	for _, product := range products {
		// O Category vai vir todo vazio
		fmt.Printf("Produto: %s Price: %.2f CategeoryID: %d Category: %s\n", product.Name, product.Price, product.Category.ID, product.Category.Name)
	}

	db.Preload("Category").Find(&products)
	for _, product := range products {
		// Após o preload, todo o category vai vir preenchido
		fmt.Printf("Produto: %s Price: %.2f CategeoryID: %d Category: %s\n", product.Name, product.Price, product.Category.ID, product.Category.Name)
	}
}
