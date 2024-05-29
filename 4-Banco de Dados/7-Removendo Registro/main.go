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

	// Create a database connection
	db_ptr, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/goexpert")

	if err != nil {
		log.Printf("Não conseguiu estabelecer conexão com o banco")
		panic(err)
	}

	defer db_ptr.Close()

	// Create a new instance of product
	product_ptr := NewProduct("Placa de Video RTX4090", 14999.99)

	// Insert the product into the database
	err = insertProduct(db_ptr, *product_ptr)
	if err != nil {
		panic(err)
	}

	// Update the local instance of the product
	product_ptr.Price = 10.99
	// Update the row of the product registration
	err = updateProduct(db_ptr, *product_ptr)
	if err != nil {
		panic(err)
	}

	// Query to find the product on the database to check the update
	product_ptr, err = selectProduct(db_ptr, product_ptr.ID)
	if err != nil {
		fmt.Printf("Não conseguiu fazer select product\n")
		panic(err)
	}
	fmt.Printf("O produto retornado do banco:\nid: %s\t name: %s\tprice: %.2f\n", product_ptr.ID, product_ptr.Name, product_ptr.Price)

	// Lista all products
	fmt.Printf("List all product on the database\n")
	products, err := selectProductsQuery(db_ptr)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Produto:\tid: %s\t name: %s\tprice: %.2f\n", p.ID, p.Name, p.Price)
	}

	// Delete the product updated
	fmt.Printf("Deleting the product with id: %s\n", product_ptr.ID)
	err = deleteProduct(db_ptr, product_ptr.ID)
	if err != nil {
		panic(err)
	}

	//list all products
	fmt.Printf("List all product on the database\n")
	products, err = selectProductsQuery(db_ptr)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Produto:\tid: %s\t name: %s\tprice: %.2f\n", p.ID, p.Name, p.Price)
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

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	// Acho isso aqui meio fraco e bem possível de dar xabú se precisa mudar o contrato do product

	//Uma outra alternativa é poder fazer o query row com um context para poder dar timeout caso necessário.
	//err = stmt.QueryRowContext( ctx, id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectProductsQuery(db *sql.DB) ([]Product, error) {
	// https://go.dev/doc/database/querying#multiple_rows

	query := "select id, name, price from products"

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Como não há a possibilidade de sofrer um SQL injection não precisamos utilizar o prepare
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {

	stmt, err := db.Prepare("delete from products where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
