package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	local_db_file string = "database.db"
)

type CotacaoUSDBRL struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {

	const port = ":8080"

	log.Printf("Server Started...\n")
	defer log.Printf("Closing server...\n")

	initializeDataBase()

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", cotacaoHandler)
	http.ListenAndServe(port, mux)
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Request: /cotacao started.")
	defer log.Println("Request: /cotacao finished.")

	cotacao_ptr, err := getCotacaoDollar()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = persistData(*cotacao_ptr)
	if err != nil {
		log.Printf("Unable to persist data: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Valor do cotacao:", cotacao_ptr.Usdbrl)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"bid":"%s"}`, cotacao_ptr.Usdbrl.Bid)))

}

func getCotacaoDollar() (*CotacaoUSDBRL, error) {

	const timeOutMilliseconds = 300
	const usd_brl_economia_url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctxWithTimeOut, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeOutMilliseconds)
	defer cancel()

	request, err := http.NewRequestWithContext(ctxWithTimeOut, http.MethodGet, usd_brl_economia_url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var cotacaoResponse CotacaoUSDBRL
	err = json.Unmarshal(body, &cotacaoResponse)
	if err != nil {
		return nil, err
	}

	return &cotacaoResponse, nil

}

func initializeDataBase() {

	db_ptr, err := sql.Open("sqlite3", local_db_file)
	if err != nil {
		log.Fatalf("Error while opening the database: %v", err)
	}
	defer db_ptr.Close()

	_, err = db_ptr.Exec(`CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY AUTOINCREMENT, bid FLOAT, timestamp DATETIME)`)
	if err != nil {
		log.Fatalf("Error while creating table: %v", err)
	}

}

func persistData(cotacao CotacaoUSDBRL) error {

	const timeOutMillis = 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*timeOutMillis)
	defer cancel()

	db_ptr, err := sql.Open("sqlite3", local_db_file)
	if err != nil {
		return err
	}
	defer db_ptr.Close()

	statement, err := db_ptr.Prepare("INSERT INTO cotacoes (bid, timestamp) VALUES (?,?);")
	if err != nil {
		return err
	}
	defer statement.Close()

	local, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return err
	}
	_, err = statement.ExecContext(ctx, cotacao.Usdbrl.Bid, time.Now().In(local))
	if err != nil {
		return err
	}

	return nil

}
