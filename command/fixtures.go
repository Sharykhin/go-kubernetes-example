package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Could not to open connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not pint database: %v", err)
	}

	fixtures, err := testfixtures.NewFolder(db, &testfixtures.MySQL{}, "fixtures")
	if err != nil {
		log.Fatalf("Could not create fixures context: %v", err)
	}

	err = fixtures.Load()
	if err != nil {
		log.Fatalf("Could not load fixtures: %v", err)
	}
}
