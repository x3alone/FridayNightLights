package web

import (
	"net/http"
)

func Web_Get(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		WebError(w, "Page not found 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		WebError(w, "Method not allowed 405", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "AsciiWeb/website/index.html")
}
