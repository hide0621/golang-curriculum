package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {

	err := godotenv.Load("infra/.env")
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", conn)
	/*
	   エラーだったらプロセスごと終了させる
	   例：2024/01/25 09:58:38 sql: unknown driver "mysql" (forgotten import?)　exit status 1
	*/
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	fmt.Println("MySQL server Successfully connected!")

	return db

}
