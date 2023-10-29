package main

import (
	"Basic-Enrollment-System/config"
	"Basic-Enrollment-System/routes"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoClient := config.DBConnect()
	routes.SetupRoutes(mongoClient)

	fmt.Println("hello db")

}
