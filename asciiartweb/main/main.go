package main

import (
	"net/http"

	serv "artweb/AsciiWeb"
)

func main() {
	http.HandleFunc("/", serv.Web_Get)
	http.ListenAndServe("localhost:8080", nil)
}
