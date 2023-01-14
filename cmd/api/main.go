package main

import (
	"github.com/gorilla/mux"
	"kravets-family-api/internal/users/handlers"
	"kravets-family-api/internal/users/services"
	"log"
	"net/http"
)

func main() {
	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(userService)

	r := mux.NewRouter()
	rV1 := r.PathPrefix("/api/v1/").Subrouter()
	rV1.HandleFunc("/user/{id}", userHandler.Show).Methods(http.MethodGet)

	rV1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route not found"))
		return
	})

	log.Fatal(http.ListenAndServe(":8000", rV1))
}
