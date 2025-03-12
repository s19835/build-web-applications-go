package data

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "wemBob-1topco-hozzoc"
	DBDbase = "web_page"
	PORT    = ":8080"
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

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	thisPage := Page{}
	fmt.Println(pageID)

	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE id=?",
		pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	if err != nil {
		log.Println("Coudn't get the page: +pageID")
		log.Println(err.Error())
	}

	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title +
		`</h1><div>` + thisPage.Content + `</div></body></html>`

	fmt.Fprintln(w, html)
}
