package main

import (
	"net/http"

	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "test:test@tcp(mysql-service:3306)/test?charset=utf8")
	if err != nil {
		log.Fatalf("Could not connetc to mysql: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping database: %v", err)
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World, Guys!"))
	})

	http.ListenAndServe(":8080", nil)
}
