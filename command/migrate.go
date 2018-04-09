package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could create migrate driver: %v", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/github.com/golang/example/outyet/migrations",
		"mysql", driver)
	if err != nil {
		log.Fatalf("Could not create database instance of migrate: %v", err)
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No changes in migrations")
		} else {
			log.Fatalf("Could not run migrate command Up: %v", err)
		}
	}

	v, _, err := m.Version()
	if err != nil {
		log.Fatalf("Get Version: %v", err)
	}

	fmt.Printf("Migrate applied. Current version is %v", v)
}
