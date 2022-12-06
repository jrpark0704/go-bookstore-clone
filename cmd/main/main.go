package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jrpark0704/go-bookstore-clone/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstore(router)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe()"localhost:9010", r)
}
