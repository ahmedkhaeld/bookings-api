// Package models describe the database in a format that go understands
package models

import "time"

// RoomRestriction is the room restriction model
type RoomRestriction struct {
	ID            int         `validate:"uuid",json:"id"`
	StartDate     time.Time   `json:"startDate"`
	EndDate       time.Time   `json:"endDate"`
	RoomID        int         `json:"roomID"`
	ReservationID int         `json:"reservationID"`
	RestrictionID int         `json:"restrictionID"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
	Room          Room        `json:"room"`
	Reservation   Reservation `json:"reservation"`
	Restriction   Restriction `json:"restriction"`
}
