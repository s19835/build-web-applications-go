package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "wemBob-1topco-hozzoc"
	DBDbase = "web_page"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func GetData() *sql.DB {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect! Check your connection.")
		log.Println(err.Error())
	} else {
		database = db
		log.Println("Database connection established.")
	}
	return database
}
