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
	fmt.Println(os.Getenv("DB_USER"))

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/test?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST")))
	if err != nil {
		log.Fatal(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"mysql", driver)
	m.Up()
}
