package main

import (
	"go-learn/main/handlers"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//Books
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	//Authors
	router.HandleFunc("/authors", handlers.GetAllAuthors).Methods("GET")
	router.HandleFunc("/authors", handlers.CreateAuthor).Methods("POST")
	router.HandleFunc("/authors/{id}", handlers.GetAuthorByID).Methods("GET")
	router.HandleFunc("/authors/{id}", handlers.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", handlers.DeleteAuthor).Methods("DELETE")

	//Category
	router.HandleFunc("/categories", handlers.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", handlers.GetCategoryByID).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", handlers.DeleteCategory).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
