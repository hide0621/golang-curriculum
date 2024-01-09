package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {

	db := initDB("example.db") // golang-curriculum/1-editing/example.db
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("table created")

	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/users", func(c echo.Context) error {

		name := c.FormValue("name")
		age, _ := strconv.Atoi(c.FormValue("age"))

		result, err := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", name, age)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		id, _ := result.LastInsertId()

		return c.JSON(http.StatusOK, &User{ID: int(id), Name: name, Age: age})

	})

	e.Start(":8080")

}
