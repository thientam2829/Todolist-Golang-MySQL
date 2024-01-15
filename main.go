package main

import (
	"golang/final/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
		log.Fatal(err)
	}
	port, exist := os.LookupEnv("PORT")
	if !exist {
		log.Println("PORT not set in .env")
		log.Fatal("PORT not set in .env")
	}

	router := routes.Init()

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal(err)
	}
}
