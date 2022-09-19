package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/driver"
	"github.com/ahmedkhaeld/bookings-api/internals/repository"
	db "github.com/ahmedkhaeld/bookings-api/internals/repository/db/mongo"
)

// Handlers  type, which gives the handler access to the configurations and the database
type Handlers struct {
	DB repository.Bookings
}

// H the handler used by the handlers
var H *Handlers

// Repo  gives the handlers a postgres database connection & functions, beside the app configurations
func Repo(conn *driver.DB) *Handlers {
	return &Handlers{
		DB: db.NewMongo(conn.Mongo),
	}
}

// New set the handlers with database and configurations
func New(h *Handlers) {
	H = h
}
