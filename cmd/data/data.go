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

func getData() *sql.DB {
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

func servePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//pageID := vars["id"]
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)
	//fmt.Println(pageID)

	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=?",
		/*pageID*/ pageGUID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		//log.Println("Coudn't get the page: +pageID")
		//log.Println(err.Error())
		log.Println("couldn't get the page!")
	}

	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title +
		`</h1><div>` + thisPage.Content + `</div></body></html>`

	fmt.Fprintln(w, html)
}

func Route() {
	getData()
	route := mux.NewRouter()
	// route.HandleFunc("/pages/{id:[0-9]+}", servePage)
	route.HandleFunc("/pages/{guid:[0-9a-zA\\-]+}", servePage)
	http.Handle("/", route)
	http.ListenAndServe(PORT, nil)
}
