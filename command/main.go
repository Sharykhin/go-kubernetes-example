package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"gopkg.in/testfixtures.v2"
	"log"
	"os"
)

func main() {
	command := flag.String("command", "", "a command (migrate|fixtures)")
	flag.Parse()
	switch *command {
	case "migrates":
		migrates()
	case "fixtures":
		fixtures()
	default:
		fmt.Println("Please specify command: migrates|fixtures")
	}
}

func fixtures() {
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

	fmt.Println("Fixtures have been loaded.")
}

func migrates() {
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
