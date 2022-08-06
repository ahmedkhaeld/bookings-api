package models

import "time"

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int         `json:"id"`
	StartDate     time.Time   `json:"start_date" :"start_date"`
	EndDate       time.Time   `json:"end_date"`
	RoomID        int         `json:"room_id"`
	ReservationID int         `json:"reservation_id"`
	RestrictionID int         `json:"restriction_id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Room          Room        `json:"room"`
	Reservation   Reservation `json:"reservation"`
	Restriction   Restriction `json:"restriction"`
}
