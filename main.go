package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/payload", WebHook).Methods("POST")

	http.ListenAndServe(":8080", r)
}
