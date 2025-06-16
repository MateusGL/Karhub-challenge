package handlers

import (
	"encoding/json"
	"karhub-beer-api/internal/dtos"
	"karhub-beer-api/internal/models"
	"karhub-beer-api/internal/services"
	"net/http"
	"sort"
)

type beerDistance struct {
	Beer     models.Beer
	Distance float64
}

func GetRecommendation(w http.ResponseWriter, r *http.Request) {
	var req dtos.RecommendationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	beers := services.GetAllBeers()
	if len(beers) == 0 {
		http.Error(w, "no beers available", http.StatusNotFound)
		return
	}

	var beerDistances []beerDistance
	for _, b := range beers {
		avg := (b.MinTemp + b.MaxTemp) / 2
		dist := abs(req.Temperature - avg)
		beerDistances = append(beerDistances, beerDistance{b, dist})
	}

	sort.SliceStable(beerDistances, func(i, j int) bool {
		if beerDistances[i].Distance == beerDistances[j].Distance {
			return beerDistances[i].Beer.Name < beerDistances[j].Beer.Name
		}
		return beerDistances[i].Distance < beerDistances[j].Distance
	})

	bestBeer := beerDistances[0].Beer

	playlist, err := services.GetSpotifyPlaylistForBeer(bestBeer.Name)
	if err != nil {
		http.Error(w, "could not fetch playlist", http.StatusInternalServerError)
		return
	}

	res := dtos.RecommendationResponse{
		BeerStyle: bestBeer.Name,
		Playlist:  playlist,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}
