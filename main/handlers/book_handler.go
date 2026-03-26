package handlers

import (
	"encoding/json"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

var books = []models.Book{}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateBook(newBook, w)

	newBook.ID = len(books) - 1

	books = append(books, newBook)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid book id"})
	}

	for idx, book := range books {
		if idx == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "book not found"})
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid book id"})
	}

	var updatedBook models.Book

	err = json.NewDecoder(r.Body).Decode(&updatedBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateBook(updatedBook, w)

	for i, book := range books {
		if book.ID == id {
			books[i].Title = updatedBook.Title
			books[i].AuthorID = updatedBook.AuthorID
			books[i].CategoryID = updatedBook.CategoryID
			books[i].Price = updatedBook.Price
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "book not found"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid book id"})
	}

	for i, book := range books {
		if book.ID == id {
			books = slices.Delete(books, i, i+1)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "book not found"})
}

func validateBook(book models.Book, w http.ResponseWriter) {
	if book.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "book title is required"})
		return
	}

	if book.AuthorID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "author is required"})
		return
	}

	if book.CategoryID < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "category is required"})
		return
	}

	if book.Price < 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "price cannot be negative"})
		return
	}
}
