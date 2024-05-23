package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Printf("Hello\n")

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	log.Println("Request inciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second): // Para vermos esse caso basta fazer um curl localhost:8080 e esperar os 5 segundos
		log.Println("Request após o time after 5 seconds")
		w.Write([]byte("hello there\n"))
	case <-ctx.Done(): // Para vermos esse caso basta fazer um curl localhost:8080 e interrompermos o curl com ctrl+c
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "request cancelada pelo cliente", 500)
	}

}
