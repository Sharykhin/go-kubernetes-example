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
	cmd := exec.Command("echo", `"${DB_PASS}"`)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("run:", out.String())
}
