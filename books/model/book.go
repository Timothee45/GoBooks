package model

type Book struct {
	ID 			int 		`json:"Id,omitempty"`
	Name		string		`json:"Name,omitempty"`
	BookType	*BookType	`json:"BookType,omitempty"`
}
type BookType struct {
	Label		string		`json:"Label,omitempty"`
}

var BookList []Book

func InitDatas() {
    BookList = append(BookList, Book{ID: 1, Name: "Le livre de la Jungle", BookType: &BookType{Label: "Aventure"}})
    BookList = append(BookList, Book{ID: 2, Name: "Oui-oui au pays des Kangourous", BookType: &BookType{Label: "Jeunesse"}})
    BookList = append(BookList, Book{ID: 3, Name: "Cendrillon", BookType: &BookType{Label: "Roman"}})
    BookList = append(BookList, Book{ID: 4, Name: "L'île au Trésor", BookType: &BookType{Label: "Aventure"}})
    BookList = append(BookList, Book{ID: 5, Name: "Titin et l'Etoile Mystérieuse", BookType: &BookType{Label: "BD"}})
}

func SelectAllBooks() []Book {
	return BookList
}

func SelectBookById(id int) Book {
	var foundBook Book

	for _, book := range BookList {
		if book.ID == id {
			foundBook = book
			break
		}
	}

	return foundBook
}

func SelectBookByTypes(bookType string) []Book {
	var filteredBooks []Book

	for _, book := range BookList {
		if book.BookType.Label == bookType {
			filteredBooks = append(filteredBooks, book)
		}
	}

	return filteredBooks
}

func InsertBook(bookName string, bookType string) Book {
    bookId := GenerateNewId()

	newBook := Book{ID: bookId, Name: bookName, BookType: &BookType{Label: bookType}}

	BookList = append(BookList, newBook)

	return newBook
}

func UpdateBookById(id int, bookName string, bookType string) Book {
	var foundBook Book

	for index, book := range BookList {
		if book.ID == id {
			foundBook = book

			foundBook.Name = bookName
			foundBook.BookType.Label = bookType

			BookList[index] = foundBook
			break
		}
	}

	return foundBook
}

func DeleteBook(id int) Book {
	var newBookList []Book
	var foundBook Book

    for _, book := range BookList {
		if book.ID != id {
			newBookList = append(newBookList, book)
		} else {
			foundBook = book
		}
	}

	BookList = newBookList

	return foundBook
}

func GenerateNewId() int {
	maxId := 0

	for _, book := range BookList {
		if book.ID > maxId {
			maxId = book.ID
		}
	}

	return maxId + 1
}
