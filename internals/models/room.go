//Package models describe the database in a format that go understands
package models

import "time"

// Room is the room model
type Room struct {
	ID        int       `validate:"uuid",json:"id,omitempty"`
	RoomName  string    `json:"roomName,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
