package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("testing\n")

	/*
	   Primeiro jeito



	*/

	/*
	   Segundo Jeito
	*/

	mux := http.NewServeMux()

	fileServe := http.FileServer(http.Dir("./public"))

	mux.Handle("/", fileServe)

	// http.ListenAndServe(":8080", mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
