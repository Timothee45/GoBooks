package handler

import (
	"github.com/gorilla/mux"
	"sort"
    "encoding/json"
	"net/http"
	"strconv"
    "books/model"
    "books/error"
    "books/data"
)

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
		json.NewEncoder(w).Encode(error.IncorrectId)
    }
}

func GetBooksByType(w http.ResponseWriter, r *http.Request) {
	arrayOrder := r.URL.Query()["order"]
	params := mux.Vars(r)

	var filteredBooks []data.Book

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
		json.NewEncoder(w).Encode(error.IncorrectId)
    }
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

    bookId, err := strconv.Atoi(params["id"])
    
    if err == nil {
    	foundBook := model.DeleteBook(bookId)

		json.NewEncoder(w).Encode(foundBook)
    } else {
		json.NewEncoder(w).Encode(error.IncorrectId)
    }
}

func OrderBooks(listBooks []data.Book, order string) []data.Book {
	selectedBooks := listBooks

	if order != "" {
		sort.Sort(NameSorter(selectedBooks))

		if order == "desc" {
			selectedBooks = ReverseArray(selectedBooks)
		}
	}

	return selectedBooks
}

// NameSorter sorts books by Name.
type NameSorter []data.Book

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

func ReverseArray(selectedBooks []data.Book) []data.Book {
	var orderedBooks []data.Book
	nbrBooks := len(selectedBooks) - 1

	for i := nbrBooks; i >= 0; i = i - 1 {
		orderedBooks = append(orderedBooks, selectedBooks[i])
	}

	return orderedBooks
}
