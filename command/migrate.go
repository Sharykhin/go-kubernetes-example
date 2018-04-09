package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Getenv("DB_USER"))

	address := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	cmd := exec.Command("goose", "-dir", "migrations", "mysql", address, "up")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migrate: ", out.String())
}
