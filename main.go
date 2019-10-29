package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	collector := make(chan *GitPayload, 100)

	r.HandleFunc("/payload", WebHook(collector)).Methods("POST")

	go func() {
		for {
			select {
			case payload := <-collector:
				err := download(payload)
				if err != nil {
					log.Printf("error downloading payload to disk: %s", err.Error())
				}
			}
		}
	}()

	http.ListenAndServe(":8080", r)
}
