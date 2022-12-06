package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jrpark0704/go-bookstore/pkg/model"
)

var NewBook model.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	books := NewBook.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

}