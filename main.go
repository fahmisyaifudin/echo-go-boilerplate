package main

import (
	"log"

	"github.com/fahmisyaifudin/echo-boilerplate/cmd"
	"github.com/fahmisyaifudin/echo-boilerplate/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
	cmd.Execute()
}
