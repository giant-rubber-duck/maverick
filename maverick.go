package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"net/http"

	"log"

	binance "github.com/adshao/go-binance/v2"
	_ "github.com/go-sql-driver/mysql"

	"rsc.io/quote"
)

type BinanceConnector struct {
	client *binance.Client
}

func (bc *BinanceConnector) GetExchangeInfoService(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	svc := bc.client.NewExchangeInfoService()
	res, err := svc.Do(ctx)
	if err != nil {
		fmt.Errorf("getting exchange info, %s", err.Error())
	}
	result := fmt.Sprintf("exchange info: %v", res)
	w.Write([]byte(result))
}

func main() {
	fmt.Println(quote.Go())

	// database
	db, err := sql.Open("mysql", "root:mysqlpw@(127.0.0.1:49153)/maverick?parseTime=true")
	if err != nil {
		fmt.Println("error when opening mysql:", err.Error())
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error when db.Ping:", err.Error())
		return
	}

	query := `
    CREATE TABLE if not exists users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println("error when db.Exec:", err.Error())
		return
	}

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	userID, err := result.LastInsertId()
	fmt.Println("last insert user id:", userID)
	var (
		id            int
		readUsername  string
		readPassword  string
		readCreatedAt time.Time
	)

	query = `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, userID).Scan(&id, &readUsername, &readPassword, &readCreatedAt)
	if err != nil {
		fmt.Println("error when db.QueryRow:", err.Error())
		return
	}
	fmt.Printf("username: %s, password: %s, created_at: %v \n", readUsername, readPassword, readCreatedAt)

	// http server
	http.HandleFunc("/", welcome)

	binanceConnector := BinanceConnector{
		client: binance.NewClient("", ""),
	}

	http.HandleFunc("/exchangeinfo", binanceConnector.GetExchangeInfoService)
	http.ListenAndServe(":80", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website!")
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
