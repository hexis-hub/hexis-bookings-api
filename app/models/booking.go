package models

import "github.com/google/uuid"

// Booking structure
type Booking struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	User  *User     `json:"user,omitempty"`
	Items []*Item   `json:"items,omitempty"`
}
