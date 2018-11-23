package models

import "github.com/google/uuid"

// Item structure
type Item struct {
	UUID       uuid.UUID `json:"uuid,omitempty"`
	Name       string    `json:"name,omitempty"`
	Type       string    `json:"type,omitempty"`
	PickedUpAt string    `json:"pickedUpAt,omitempty"`
	ReturnAt   string    `json:"returnAt,omitempty"`
}
