package repository

import (
	"github.com/ahmedkhaeld/bookings-api/internals/models"
)

//Bookings defines the database methods to be implemented
type Bookings interface {
	FetchAllRooms() ([]*models.Room, error)
	FetchRoom(id int) (models.Room, error)
	UpdateRoom(r models.Room) error
	InsertRoom(r models.Room) error
	DeleteRoom(r models.Room) error
}
