// Package models describe the database in a format that go understands
package models

import "time"

// Restriction is the restriction model
type Restriction struct {
	ID              int       `validate:"uuid",json:"id,omitempty"`
	RestrictionName string    `json:"restrictionName,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
