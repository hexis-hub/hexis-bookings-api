package models

import "github.com/google/uuid"

// User structure
type User struct {
	UUID  uuid.UUID `json:"uuid,omitempty"`
	Name  string    `json:"name,omitempty"`
	Email string    `json:"email,omitempty"`
}
