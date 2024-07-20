package web

import (
	"html/template"
	"net/http"
)

func Web_Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed: error 400", http.StatusBadRequest)
		return
	}

	Data, err := Collect_info(r)
	if err != nil {
		http.Error(w, "Character Not Allowed: error 501", http.StatusNotImplemented)
		return
	}

	t, err := template.ParseFiles("AsciiWeb/post.html")
	if err != nil {
		http.Error(w, "File Parsing Failure: error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Data)
	if err != nil {
		http.Error(w, "Execution Failure: error 500", http.StatusInternalServerError)
		return
	}
}
