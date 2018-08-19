package main

import ("github.com/gorilla/mux"
	"log"
	"net/http")

func main() {

	InitBooks()

	get("/api/books", GetAllBooks)
	get("/api/books/{id}", GetBookByID)
	post("/api/books", CreateBook)
	put("/api/books/{id}", UpdateBook)
	delete("/api/books/{id}", DeleteBook)

	startServer(":8000")
}

// ControllerFunc type
type ControllerFunc func(http.ResponseWriter, *http.Request)

var router = mux.NewRouter()

func get(path string, cf ControllerFunc) {
	router.HandleFunc(path, cf).Methods("GET")
}

func post(path string, cf ControllerFunc) {
	router.HandleFunc(path, cf).Methods("POST")
}

func put(path string, cf ControllerFunc) {
	router.HandleFunc(path, cf).Methods("PUT")
}

func delete(path string, cf ControllerFunc) {
	router.HandleFunc(path, cf).Methods("DELETE")
}

func startServer(port string) {
	log.Fatal(http.ListenAndServe(port, router))
}
