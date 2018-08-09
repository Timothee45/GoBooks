package handler

import (
	"github.com/gorilla/mux"
	"sort"
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
	arrayOrder := r.URL.Query()["order"]

	selectedBooks := model.SelectAllBooks()

	if len(arrayOrder) != 0 {
		selectedBooks = OrderBooks(selectedBooks, arrayOrder[0])
	}

	json.NewEncoder(w).Encode(selectedBooks)
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
	arrayOrder := r.URL.Query()["order"]
	params := mux.Vars(r)
	
	var filteredBooks []model.Book

	filteredBooks = model.SelectBookByTypes(params["type"])

	if len(arrayOrder) != 0 {
		filteredBooks = OrderBooks(filteredBooks, arrayOrder[0])
	}

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

func OrderBooks(listBooks []model.Book, order string) []model.Book {
	selectedBooks := listBooks

	if order != "" {
		sort.Sort(NameSorter(selectedBooks))

		if order == "desc" {
			selectedBooks = ReverseArray(selectedBooks)
		}
	}

	return selectedBooks
}

// NameSorter sorts planets by name.
type NameSorter []model.Book

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func ReverseArray(selectedBooks []model.Book) []model.Book {
	var orderedBooks []model.Book
	nbrBooks := len(selectedBooks) - 1

	for i := nbrBooks; i >= 0; i = i - 1 {
		orderedBooks = append(orderedBooks, selectedBooks[i])
	}

	return orderedBooks
}
