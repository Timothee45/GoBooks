package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"books/handler"
	"books/error"
    "encoding/json"
    "fmt"
)

func InitRoutes() {
    router := mux.NewRouter()

	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/book", handler.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", handler.ModifyBook).Methods("PUT")
	router.HandleFunc("/book/{id}", handler.RemoveBook).Methods("DELETE")

	router.HandleFunc("/books/type/{type}", handler.GetBooksByType).Methods("GET")

	router.Use(Middleware)

    log.Fatal(http.ListenAndServe(":8000", router))
}

func Middleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	token := r.Header.Get("user")

    	if token != "" {
    		resp, err := http.Get("http://localhost:8080/?user=" + token)

    		fmt.Println(resp.Body)

    		if err == nil {
    			log.Println("middleware", r.URL)
        		h.ServeHTTP(w, r)
    		}
    	} else {
    		json.NewEncoder(w).Encode(error.BadToken)
    	}
    })
}
