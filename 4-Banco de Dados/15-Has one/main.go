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
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category     // Aqui todo produto tem uma categoria
	SerialNumber SerialNumber // Aqui todo produto tem que tem um numero serial. Mas precisamos fazer ele ser unico
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	fmt.Printf("Hello there\n")

	dsn := "root:password@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// Criando uma nova categoria
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	// Criando um novo produto nessa categoria
	new_product := Product{
		Name:       "Polystation",
		Price:      49.90,
		CategoryID: category.ID, // Aqui estabelecemos a relação que o produto belongs to uma category
	}
	db.Create(&new_product)

	//Criando um novo serial number para ligarmos com o product
	serial_number := SerialNumber{
		Number:    "123456",
		ProductID: new_product.ID,
		/*
		   Aqui estabelecemos a relação que o serial number belongs to new_product ou
		   tambem podemos falar que o new_product has one SerialNumber
		*/
	}
	db.Create(&serial_number)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	for _, value := range products {
		fmt.Printf("Product: %s, Category: %s, SerialNumber: %s\n", value.Name, value.Category.Name, value.SerialNumber.Number)
	}
}
