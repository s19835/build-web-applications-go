package web

import (
	"net/http"
)

func FileServe() {
	http.ListenAndServe(PORT,
		http.FileServer(http.Dir("/var/www")),
	)
}
