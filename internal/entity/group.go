package entity

import "github.com/google/uuid"

type Group struct {
	ID          uuid.UUID `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
