package main

import (
	"database/sql"
	"fmt"
	"time"

	"net/http"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	// database
	db, err := sql.Open("mysql", "root:1234@(127.0.0.1:3306)/new_schema?parseTime=true")
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})
	http.ListenAndServe(":80", logging(foo))
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
