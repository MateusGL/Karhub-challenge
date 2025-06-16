package routes

import (
	"log"
	"net/http"

	"karhub-beer-api/internal/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

	r.HandleFunc("/beers", handlers.ListBeers).Methods("GET")
	r.HandleFunc("/beers", handlers.CreateBeer).Methods("POST")
	r.HandleFunc("/beers/{id}", handlers.UpdateBeer).Methods("PUT")
	r.HandleFunc("/beers/{id}", handlers.DeleteBeer).Methods("DELETE")

	r.HandleFunc("/recommendation", handlers.GetRecommendation).Methods("POST")

	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
