package handlers

import (
	"github.com/ahmedkhaeld/bookings-api/internals/driver"
	"github.com/ahmedkhaeld/bookings-api/internals/repository"
	db "github.com/ahmedkhaeld/bookings-api/internals/repository/db/postgres"
)

// Handlers   has access the Bookings' db functions
type Handlers struct {
	DB repository.Bookings
}

// H the handler used by the handlers
var H *Handlers

// Repo  gives the handlers a database type, its methods implement the bookings interface
func Repo(conn *driver.DB) *Handlers {
	return &Handlers{
		DB: db.PostgresRepo(conn.SQL),
	}
}

// New set the handlers with database and configurations
func New(r *Handlers) {
	H = r
}
