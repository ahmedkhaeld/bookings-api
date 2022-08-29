//Package models describe the database in a format that go understands
package models

import "time"

// Room is the room model
type Room struct {
	ID        int       `json:"id,omitempty"`
	RoomName  string    `validate:"required",json:"roomName,omitempty"`
	Price     int       `validate:"required",json:"price"`
	Rate      int       `validate:"required",json:"rate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
