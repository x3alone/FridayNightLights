package main

import (
	"net/http"

	serv "artweb/AsciiWeb"
)

func main() {
	http.HandleFunc("/", serv.Web_Get)
	http.HandleFunc("/ascii-art", serv.Web_Post)
	http.ListenAndServe("localhost:8080", nil)
}
