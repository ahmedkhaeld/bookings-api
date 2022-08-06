package models

import "time"

// Restriction is the restriction model
type Restriction struct {
	ID              int       `json:"id,omitempty"`
	RestrictionName string    `json:"restriction_name,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
