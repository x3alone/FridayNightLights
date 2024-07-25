package web

import (
	"html/template"
	"net/http"
)

func WebError(w http.ResponseWriter, msj string, code int) {

	w.WriteHeader(code)

	t, err := template.ParseFiles("AsciiWeb/website/error.html")

	if err != nil {
		http.Error(w, "File Parsing Failure: error 500", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, msj)
	
	if err != nil {
		http.Error(w, "Execution Failure: error 500", http.StatusInternalServerError)
		return
	}
}
