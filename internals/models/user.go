package models

import "time"

// User is the user model
type User struct {
	ID          int       `json:"id,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
	AccessLevel int       `json:"access_level,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
