package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patito/git-notifier/notifier"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", notifier.Index)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
