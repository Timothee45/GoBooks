package handler

import (
	"github.com/gorilla/mux"
    "encoding/json"
	"net/http"
	"strconv"
    "books/model"
)

type Error struct {
	Code 		int
	Label		string
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(model.SelectAllBooks())
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

    bookId, err := strconv.Atoi(params["id"])

    if err == nil {
		foundBook := model.SelectBookById(bookId)

		json.NewEncoder(w).Encode(foundBook)
    } else {
		json.NewEncoder(w).Encode(Error{Code: 415, Label: "Incorrect id"})
    }
}

func GetBooksByType(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var filteredBooks []model.Book

	filteredBooks = model.SelectBookByTypes(params["type"])

	json.NewEncoder(w).Encode(filteredBooks)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookName := r.FormValue("Name")
	bookType := r.FormValue("BookType")

	newBook := model.InsertBook(bookName, bookType)

	json.NewEncoder(w).Encode(newBook)
}

func ModifyBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookName := r.FormValue("Name")
	bookType := r.FormValue("BookType")

    bookId, err := strconv.Atoi(params["id"])
    
    if err == nil {
		foundBook := model.UpdateBookById(bookId, bookName, bookType)

		json.NewEncoder(w).Encode(foundBook)
    } else {
		json.NewEncoder(w).Encode(Error{Code: 415, Label: "Incorrect id"})
    }
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

    bookId, err := strconv.Atoi(params["id"])
    
    if err == nil {
    	foundBook := model.DeleteBook(bookId)

		json.NewEncoder(w).Encode(foundBook)
    } else {
		json.NewEncoder(w).Encode(Error{Code: 415, Label: "Incorrect id"})
    }
}
