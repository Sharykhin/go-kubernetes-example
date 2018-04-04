package main

import (
	"net/http"

	"database/sql"
	"log"

	"encoding/json"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

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

type user struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		var u user
		row := db.QueryRow("SELECT * FROM test LIMIT 1")
		err := row.Scan(&u.Id, &u.Text)
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
		w.Write(res)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World."))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong 2"))
	})

	http.ListenAndServe(":8080", nil)
}
