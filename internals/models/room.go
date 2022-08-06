package models

import "time"

// Room is the room model
type Room struct {
	ID        int       `json:"id,omitempty"`
	RoomName  string    `json:"room_name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
