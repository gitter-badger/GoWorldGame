package main

import (
	"fmt"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	mux.HandleFunc("/register", RegisterNewUser)
	
}

func RegisterNewUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Register new User!")
}
