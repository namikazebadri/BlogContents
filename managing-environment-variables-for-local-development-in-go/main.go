package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file not found")
	}

	envVars, err := godotenv.Read()

	fmt.Printf("Hello %s %s!", envVars["FIRST_NAME"], envVars["LAST_NAME"])
}
