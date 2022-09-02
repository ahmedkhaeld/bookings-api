package repository

import (
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"time"
)

//Bookings defines the database methods to be implemented
type Bookings interface {
	FetchAllRooms() ([]*models.Room, error)
	FetchRoom(id int) (models.Room, error)
	UpdateRoom(r models.Room) error
	InsertRoom(r models.Room) error
	DeleteRoom(r models.Room) error
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
}
