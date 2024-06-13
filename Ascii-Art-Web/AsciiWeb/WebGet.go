package web

import (
	"net/http"
)

func Web_Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page not found: error 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed: error 405", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "AsciiWeb/index.html")
}
