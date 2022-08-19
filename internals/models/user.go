// Package models describe the database in a format that go understands
package models

import (
	"time"
)

// User is the user model
type User struct {
	ID          int       `validate:"uuid",json:"id,omitempty"`
	FirstName   string    `validate:"required",json:"firstName,omitempty"`
	LastName    string    `validate:"required",json:"lastName,omitempty"`
	Email       string    `validate:"required,email",json:"email,omitempty"`
	Password    string    `validate:"required,gte=10",json:"password,omitempty"`
	AccessLevel int       `json:"accessLevel,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
