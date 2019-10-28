package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/payload", WebHook).Methods("POST")

	http.ListenAndServe(":8080", r)
}

// WebHook listens for Github repo updates
func WebHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Print(string(body))
	w.WriteHeader(http.StatusOK)

}
