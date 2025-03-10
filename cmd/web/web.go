package web

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"

	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}

	http.ServeFile(w, r, fileName)
}

func Serve() {
	router := mux.NewRouter()
	router.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	http.Handle("/", router)

	http.ListenAndServe(PORT, nil)
}
