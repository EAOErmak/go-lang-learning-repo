package handlers

import (
	"encoding/json"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

//Author
//Authors
//author
//authors

var authors = []models.Author{}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	var newAuthor models.Author

	err := json.NewDecoder(r.Body).Decode(&newAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateAuthor(newAuthor, w)

	newAuthor.ID = len(authors) - 1

	authors = append(authors, newAuthor)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAuthor)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
	}

	for idx, author := range authors {
		if idx == id {
			json.NewEncoder(w).Encode(author)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "author not found"})
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
	}

	var updatedAuthor models.Author

	err = json.NewDecoder(r.Body).Decode(&updatedAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateAuthor(updatedAuthor, w)

	for i, author := range authors {
		if author.ID == id {
			authors[i].Name = updatedAuthor.Name
			json.NewEncoder(w).Encode(authors[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "author not found"})
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
	}

	for i, author := range authors {
		if author.ID == id {
			authors = slices.Delete(authors, i, i+1)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "author not found"})
}

func validateAuthor(author models.Author, w http.ResponseWriter) {
	if author.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "author title is required"})
		return
	}
}
