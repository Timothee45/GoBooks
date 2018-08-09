package data

import 	(
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"log"
)

var BOOKS_DATA_PATH = "data/books.json"

type Book struct {
	ID 			int 		`json:"Id,omitempty"`
	Name		string		`json:"Name,omitempty"`
	BookType	*BookType	`json:"BookType,omitempty"`
}

type BookType struct {
	Label		string		`json:"Label,omitempty"`
}

func ReadDatas() []Book {
	jsonFile, err := os.Open(BOOKS_DATA_PATH)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// we create the slice of Books
	var BookList []Book

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &BookList)

	return BookList
}

func WriteData(bookList []Book) {
	b, err := json.Marshal(bookList)

	if err != nil {
		log.Fatal(err)
	}

	// the WriteFile method returns an error if unsuccessful
	err = ioutil.WriteFile(BOOKS_DATA_PATH, b, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}
}
