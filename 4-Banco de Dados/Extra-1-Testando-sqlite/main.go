package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal("Unable to connect to database")
		return
	}
	defer db.Close()

	if _, err := db.Exec(`
	  CREATE TABLE IF NOT EXISTS quotations (
	  id varchar(255) PRIMARY KEY,
	  bid varchar(255)
	  );
	`); err != nil {
		log.Fatal("Unable to create table")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Millisecond))
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO quotations VALUES(?,?);")
	if err != nil {
		log.Fatal("Unable to create prepare statement")
		return
	}

	_, err = stmt.ExecContext(ctx, "123", "5.2")
	if err != nil {
		log.Fatal("Unable to execute create")
	}
}
