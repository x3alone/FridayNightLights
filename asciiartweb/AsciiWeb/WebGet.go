package web

import (
	"html/template"
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

	t, err := template.ParseFiles("./static/testget.html")
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, "Method Not Allowed: error 500", http.StatusInternalServerError)
		return
	}
}
