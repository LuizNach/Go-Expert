package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
No contexto de bancos de dados, as diferenças entre lock otimista e lock pessimista estão relacionadas à
forma como os sistemas lidam com o controle de concorrência.

* Lock pessimista: Nesse modelo, o sistema assume que haverá conflitos de acesso aos dados e, portanto,
bloqueia recursos de forma mais agressiva, impedindo que outros usuários acessem ou modifiquem os dados
enquanto uma transação está em andamento. Isso pode levar a potenciais bloqueios e diminuição de desempenho
em situações de alta concorrência.

=> Faz um lock na linha da table;

* Lock otimista: Já no lock otimista, o sistema parte do pressuposto de que conflitos de acesso aos dados
são menos prováveis. Dessa forma, não são aplicados bloqueios preventivos, permitindo que múltiplos usuários
acessem e modifiquem os dados simultaneamente. Antes de efetivar as alterações, o sistema verifica se houve
alguma interferência, podendo então realizar ações de correção, se necessário.

=> Versiona quando alguém está fazendo alguma alteraçpao no sistema.

Essas abordagens têm impacto direto na eficiência, escalabilidade e consistência dos sistemas de banco de
dados, sendo importante escolher a estratégia mais adequada de acordo com as necessidades e características
do projeto.
*/

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"` // Aqui estabelecemos que havera uma tabela que faz a relação many to many
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"` // Precisamos declarar aqui tambem que a relacao é many to many nesse ponto
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

	// Criando duas novas categorias
	category := Category{Name: "Eletronics"}
	db.Create(&category)

	category2 := Category{Name: "VideoGames"}
	db.Create(&category)

	// Criando um novo produto com as duas categorias
	new_product := Product{
		Name:       "Polystation",
		Price:      49.90,
		Categories: []Category{category, category2}, // CategoryID não faz mais sentido passamos as novas categorias num slice de Category
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
		Name:       "Multimetro",
		Price:      199.99,
		Categories: []Category{category},
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
		fmt.Printf("Product: %s, SerialNumber: %s\n", value.Name, value.SerialNumber.Number)
	}

	var categories []Category
	// Listando todas as categorias encontradas no banco com seus produtos
	err = db.Debug().Model(&Category{}).Preload("Products.Categories").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		fmt.Printf("Deu erro buscar por categorias e popular os produtos das categorias\n")
		panic(err)
	}

	for _, value := range categories {
		fmt.Printf("\nCategoria: %s\n", value.Name)

		for _, product := range value.Products {
			fmt.Printf("Name: %s, Price: %.2f, SerialNumber: %s\n", product.Name, product.Price, product.SerialNumber.Number)
		}
	}

	// Aqui começamos a fazer um lock para fazer um update no banco.
	// Isso é um Lock Pessimista: travamos o dado e ninguém mais pode atualizar enquanto a operação não terminar
	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	// Aqui buscamos um unico elemento de category e damos lock nele
	// A query podemos ver que vai ser:
	// SELECT * FROM `categories` WHERE `categories`.`id` = 1 ORDER BY `categories`.`id` LIMIT 1 FOR UPDATE
	// Esse FOR UPDATE mostra que ele estará "lockado" para qualquer atualização

	if err != nil {
		panic(err)
	}

	c.Name = "Eletronics"
	tx.Debug().Save(&c)
	tx.Commit()
}
