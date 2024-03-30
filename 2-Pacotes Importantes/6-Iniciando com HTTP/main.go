package main

import "net/http"

func main() {
	//http.ListenAndServe("8080", nil) // com uma linha ja temos um servidor http

	//multiplexer :=

	// Quando eu quiser conectar uma funcao a um path, utilizo a funcao HandleFunc
	http.HandleFunc("/", Root)

	// Tambem é possivel attachar uma funcao anonima
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":true}`))
	})

	// iniciamos o servidor com Listen and Serve
	http.ListenAndServe(":8080", nil)
}

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from root\n"))
}
