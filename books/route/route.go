package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"books/handler"
)

func InitRoutes() {
    router := mux.NewRouter()

	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/book", handler.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", handler.ModifyBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handler.RemoveBook).Methods("DELETE")

	router.HandleFunc("/books/type/{type}", handler.GetBooksByType).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", router))
}
