package handlers

import (
	"encoding/json"
	"go-learn/main/models"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

var categories = []models.Category{}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	var newCategory models.Category

	err := json.NewDecoder(r.Body).Decode(&newCategory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateCategory(newCategory, w)

	newCategory.ID = len(categories) - 1

	categories = append(categories, newCategory)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid category id"})
	}

	for idx, category := range categories {
		if idx == id {
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "category not found"})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid category id"})
	}

	var updatedCategory models.Category

	err = json.NewDecoder(r.Body).Decode(&updatedCategory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
		return
	}

	validateCategory(updatedCategory, w)

	for i, category := range categories {
		if category.ID == id {
			categories[i].Name = updatedCategory.Name
			json.NewEncoder(w).Encode(categories[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "category not found"})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid category id"})
	}

	for i, category := range categories {
		if category.ID == id {
			categories = slices.Delete(categories, i, i+1)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "category not found"})
}

func validateCategory(category models.Category, w http.ResponseWriter) {
	if category.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "category title is required"})
		return
	}
}
