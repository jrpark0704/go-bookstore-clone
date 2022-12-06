package routes

import (
	"github.com/gorilla/mux"
	"github.com/jrpark0704/go-bookstore-clone/pkg/controllers"
)

func RegisterBookstore(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
}
