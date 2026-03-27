package handlers

import (
	"encoding/json"
	"errors"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

var authors = []models.Author{}
var nextAuthorID = 1

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newAuthor models.Author

	err := json.NewDecoder(r.Body).Decode(&newAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	err = validateAuthor(newAuthor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	newAuthor.ID = nextAuthorID
	nextAuthorID++

	authors = append(authors, newAuthor)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAuthor)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
		return
	}

	for _, author := range authors {
		if author.ID == id {
			json.NewEncoder(w).Encode(author)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "author not found"})
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
		return
	}

	var updatedAuthor models.Author

	err = json.NewDecoder(r.Body).Decode(&updatedAuthor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	err = validateAuthor(updatedAuthor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

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
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid author id"})
		return
	}

	for i, author := range authors {
		if author.ID == id {
			authors = slices.Delete(authors, i, i+1)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "author deleted"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "author not found"})
}

func validateAuthor(author models.Author) error {
	if author.Name == "" {
		return errors.New("author name is required")
	}
	return nil
}

func AuthorExists(authorID int) bool {
	for _, author := range authors {
		if author.ID == authorID {
			return true
		}
	}
	return false
}
