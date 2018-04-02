package main

import (
	"net/http"

	"database/sql"
	"log"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB
var env = os.Getenv("APP_ENV")

func init() {
	if env == "testing" {
		var err error
		var address = os.Getenv("MYSQL_ADDRESS")
		db, err = sql.Open("mysql", address)
		if err != nil {
			log.Fatalf("Could not connetc to mysql: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Could not ping database: %v", err)
		}
	}
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var u user
		row := db.QueryRow("SELECT * FROM users LIMIT 1")
		err := row.Scan(&u.Id, &u.Name)
		if err != nil {
			log.Fatalf("Could not scan user: %v\n", err)
		}
		res, err := json.Marshal(&u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8080", nil)
}
