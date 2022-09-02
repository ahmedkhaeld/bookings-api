package main

import (
	handlers "github.com/ahmedkhaeld/bookings-api/internals/handlers/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	//api.bookings.v1.restapi.org  hostName

	mux.Get("/rooms", handlers.H.GetRooms)
	mux.Post("/rooms", handlers.H.PostRoom)

	mux.Get("/rooms/{room-id}", handlers.H.GetRoom)
	mux.Put("/rooms/{room-id}", handlers.H.PutRoom)
	mux.Delete("/rooms/{room-id}", handlers.H.DeleteRoom)

	mux.Post("/rooms/{room-id}/check-availability", handlers.H.PostCheckSingleRoomAvailability)
	mux.Post("/rooms/check-availability", handlers.H.PostCheckAllRoomsAvailability)

	return mux
}
