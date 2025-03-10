package server

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}

func Start() {
	http.HandleFunc("/", testHandler)

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}

	// router := mux.NewRouter()
	// router.HandleFunc("/test", testHandler)

	// http.Handle("/", router)
	// fmt.Println("Everything is set up!")
}
