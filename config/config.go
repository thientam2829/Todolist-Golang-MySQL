package config

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Database() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
		log.Fatal(err)
	}
	user, exist := os.LookupEnv("DB_USER")
	if !exist {
		log.Println("DB_USER not set in .env")
		log.Fatal("DB_USER not set in .env")
	}

	pass, exist := os.LookupEnv("DB_PASS")
	if !exist {
		log.Println("DB_PASS not set in .env")
		log.Fatal("DB_PASS not set in .env")
	}

	host, exist := os.LookupEnv("DB_HOST")
	if !exist {
		log.Println("DB_HOST not set in .env")
		log.Fatal("DB_HOST not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:3307)/?charset=utf8&parseTime=True", user, pass, host)
	database, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Fatal(err)
	}
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE IF NOT EXISTS todolist`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = database.Exec(`USE todolist`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	return database
}
