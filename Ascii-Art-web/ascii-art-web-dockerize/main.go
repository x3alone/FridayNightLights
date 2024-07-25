package main

import (
	"fmt"
	"log"
	"net/http"

	serv "artweb/AsciiWeb"
)

func main() {
	port := ":8080"

	/*Handers*/
	http.HandleFunc("/", serv.Web_Get)
	http.HandleFunc("/ascii-art", serv.Web_Post)
	http.Handle("/ascii/", http.StripPrefix("/ascii/", http.FileServer(http.Dir("./AsciiWeb/style/"))))

	log.Println("Serving files on " + port + "...")
	log.Println("http://localhost" + port + "/")

	/*Serve*/
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
