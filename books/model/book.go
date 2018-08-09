package model

import 	(
	"books/data"
)

type Book struct {
	ID 			int 		`json:"Id,omitempty"`
	Name		string		`json:"Name,omitempty"`
	BookType	*BookType	`json:"BookType,omitempty"`
}

type BookType struct {
	Label		string		`json:"Label,omitempty"`
}

func SelectAllBooks() []data.Book {
	BookList := data.ReadDatas()

	return BookList
}

func SelectBookById(id int) data.Book {
	BookList := data.ReadDatas()
	var foundBook data.Book

	for _, book := range BookList {
		if book.ID == id {
			foundBook = book
			break
		}
	}

	return foundBook
}

func SelectBookByTypes(bookType string) []data.Book {
	BookList := data.ReadDatas()
	var filteredBooks []data.Book

	for _, book := range BookList {
		if book.BookType.Label == bookType {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks
}

func InsertBook(bookName string, bookType string) data.Book {
	BookList := data.ReadDatas()
    bookId := GenerateNewId(BookList)

	newBook := data.Book{ID: bookId, Name: bookName, BookType: &data.BookType{Label: bookType}}

	BookList = append(BookList, newBook)

	data.WriteData(BookList)

	return newBook
}

func UpdateBookById(id int, bookName string, bookType string) data.Book {
	BookList := data.ReadDatas()
	var foundBook data.Book

	for index, book := range BookList {
		if book.ID == id {
			foundBook = book

			foundBook.Name = bookName
			foundBook.BookType.Label = bookType

			BookList[index] = foundBook
			break
		}
	}

	data.WriteData(BookList)

	return foundBook
}

func DeleteBook(id int) data.Book {
	BookList := data.ReadDatas()
	var newBookList []data.Book
	var foundBook data.Book

    for _, book := range BookList {
		if book.ID != id {
			newBookList = append(newBookList, book)
		} else {
			foundBook = book
		}
	}

	data.WriteData(newBookList)

	return foundBook
}

func GenerateNewId(bookList []data.Book) int {
	maxId := 0

	for _, book := range bookList {
		if book.ID > maxId {
			maxId = book.ID
		}
	}

	return maxId + 1
}
