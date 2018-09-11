package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/home/", home)

	server := &http.Server{
		Addr: "127.0.0.1:8080",
	}
	server.ListenAndServe()
}
