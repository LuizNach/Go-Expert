package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	fmt.Printf("Hello there\n")

	db_ptr, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goexpert")

	if err != nil {
		log.Printf("Não conseguiu estabelecer conexão com o banco")
		panic(err)
	}

	defer db_ptr.Close()

	product_ptr := NewProduct("Placa de Video RTX4090", 14999.99)

	err = insertProduct(db_ptr, *product_ptr)
	if err != nil {
		panic(err)
	}

	product_ptr.Price = 10.99
	err = updateProduct(db_ptr, *product_ptr)
	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, product Product) error {
	// prepare statement
	statement, err := db.Prepare("insert into products(id, name, price) values (?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	//
	_, err = statement.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product Product) error {

	//prepare statement
	statement, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	//
	_, err = statement.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}
