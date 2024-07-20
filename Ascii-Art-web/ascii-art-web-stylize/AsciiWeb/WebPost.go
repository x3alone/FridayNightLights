package web

import (
	"html/template"
	"net/http"
)

func Web_Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		WebError(w, "Bad request 400", http.StatusBadRequest)
		return
	}

	Data, err := Collect_info(r)
	if err != nil {
		WebError(w, "Bad request 400", http.StatusBadRequest)
		return
	}
	t, err := template.ParseFiles("AsciiWeb/website/post.html")
	if err != nil {
		WebError(w, "File Parsing Failure: error 500", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, Data)
	if err != nil {
		WebError(w, "Execution Failure: error 500", http.StatusInternalServerError)
		return
	}
}
