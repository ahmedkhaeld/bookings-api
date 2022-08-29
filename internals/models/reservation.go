// Package models describe the database in a format that go understands
package models

import "time"

// Reservation is the reservation model
type Reservation struct {
	ID        int       `validate:"uuid",json:"id,omitempty"`
	FirstName string    `validate:"required",json:"firstName,omitempty"`
	LastName  string    `validate:"required",json:"lastName,omitempty"`
	Email     string    `validate:"required",json:"email,omitempty"`
	Phone     string    `validate:"required",json:"phone,omitempty"`
	StartDate time.Time `validate:"required",json:"startDate"`
	EndDate   time.Time `validate:"required",json:"endDate"`
	RoomID    int       `json:"roomID,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Room      Room      `json:"room"`
	Processed int       `json:"processed,omitempty"`
}
