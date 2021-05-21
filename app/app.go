package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	mux.HandleFunc("/customers", customers)

	log.Fatal(http.ListenAndServe("localhost:7001", mux))
}
