package main


import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)


// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	}
	
	// Author Struct (Model)
	type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	}
	
	// Init books var as slice Book struct
	var books []Book
	


// InitBooks initlaize book data
func InitBooks() {
	books = append(books,
		Book{ID: "1",
			Isbn: "48743", Title: "Book One",
			Author: &Author{
				Firstname: "Lem",
				Lastname:  "Adane",
			},
		})

	books = append(books,
		Book{ID: "2",
			Isbn: "48746", Title: "Book Two",
			Author: &Author{
				Firstname: "LJ",
				Lastname:  "Adane",
			},
		})
}


// GetAllBooks - get all books
func GetAllBooks(response http.ResponseWriter, request *http.Request) {
	toJSON(response, books)
}

// GetBookByID - get a book by id
func GetBookByID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range books {
		if item.ID == params["id"] {
			toJSON(response, item)
			return
		}
	}
	toJSON(response, &Book{})
}

// CreateBook - add a new book
func CreateBook(response http.ResponseWriter, request *http.Request) {
	var book Book
	_ = fromJSON(request.Body, &book)
	book.ID = strconv.Itoa(rand.Intn(10000000)) // generate random id
	books = append(books, book)
	toJSON(response, book)
}

// DeleteBook - updates a book
func DeleteBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	toJSON(response, books)
}

// UpdateBook - delete a book
func UpdateBook(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = fromJSON(request.Body, &book)
			book.ID = params["id"]
			books = append(books, book)
			toJSON(response, book)
			return
		}
	}
	toJSON(response, books)
}

// Any type
type Any interface{}

func toJSON(response http.ResponseWriter, any Any) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(any)
}

func fromJSON(requestBody io.Reader, any Any) {
	json.NewDecoder(requestBody).Decode(any)
}