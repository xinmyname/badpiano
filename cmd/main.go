package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Writing data.lua...")

	file, err := os.Create("data.lua")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("data={}\n")

	fmt.Println("OK.")
}
