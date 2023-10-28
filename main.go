package main

import (
	"Basic-Enrollment-System/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.DBConnect()
	fmt.Printf("hello db")
}
