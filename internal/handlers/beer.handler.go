package handlers

import (
	"encoding/json"
	"karhub-beer-api/internal/models"
	"karhub-beer-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListBeers(w http.ResponseWriter, r *http.Request) {
	beers := services.GetAllBeers()
	json.NewEncoder(w).Encode(beers)
}

func CreateBeer(w http.ResponseWriter, r *http.Request) {
	var beer models.Beer
	json.NewDecoder(r.Body).Decode(&beer)
	saved := services.SaveBeer(beer)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(saved)
}

func UpdateBeer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var beer models.Beer
	json.NewDecoder(r.Body).Decode(&beer)
	updated, err := services.UpdateBeer(id, beer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func DeleteBeer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := services.DeleteBeer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
