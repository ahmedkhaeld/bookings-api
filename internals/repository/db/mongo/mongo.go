package db

import (
	"github.com/ahmedkhaeld/bookings-api/internals/models"
	"github.com/ahmedkhaeld/bookings-api/internals/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Mongo struct {
	DB *mongo.Client
}

func (m Mongo) FetchAllRooms() ([]*models.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) FetchRoom(id int) (models.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) UpdateRoom(r models.Room) error {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) InsertRoom(r models.Room) error {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) DeleteRoom(r models.Room) error {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (m Mongo) SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongo(conn *mongo.Client) repository.Bookings {
	return &Mongo{
		DB: conn,
	}
}
