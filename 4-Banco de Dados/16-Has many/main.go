package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product // Aqui estamos dizendo que uma category pode ter muitos products
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
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
	}
	db.Create(&serial_number)

	// Criando um segundo produto
	new_product = Product{
		Name:       "Atari",
		Price:      199.99,
		CategoryID: category.ID,
	}
	db.Create(&new_product)

	// Para o segundo produto estabelecemos um novo serial number
	serial_number = SerialNumber{
		Number:    "123457",
		ProductID: new_product.ID,
	}
	db.Create(&serial_number)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	// Listando todos os produtos criados no banco
	for _, value := range products {
		fmt.Printf("Product: %s, Category: %s, SerialNumber: %s\n", value.Name, value.Category.Name, value.SerialNumber.Number)
	}

	var categories []Category
	// Listando todas as categorias encontradas no banco com seus produtos
	// err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	// Não foi ensinado isso no curso mas podemos fazer o preload encadeado
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		fmt.Printf("Deu erro buscar por categorias e popular os produtos das categorias\n")
		panic("Deu ruim")
	}

	for _, value := range categories {
		fmt.Printf("\nCategoria: %s\n", value.Name)

		for _, product := range value.Products {
			fmt.Printf("Name: %s, Price: %.2f, CategoryName: %s, CategoryID: %d, SerialNumber: %s\n", product.Name, product.Price, product.Category.Name, product.CategoryID, product.SerialNumber.Number)
		}
	}
}
