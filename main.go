package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Need by default driver
	"log"
	"net/http"
	"os"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/test?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatalf("Could not connetc to mysql: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}
}

type user struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var u user
		row := db.QueryRow("SELECT * FROM test LIMIT 1")
		err := row.Scan(&u.ID, &u.Text)
		switch {
		case err == sql.ErrNoRows:
		case err != nil:
			{
				log.Fatalf("Could not scan user: %v\n", err)
			}
		}
		res, err := json.Marshal(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res) // nolint: errcheck
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!!!")) // nolint: errcheck
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong")) // nolint: errcheck
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
