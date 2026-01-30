package models

import "time"

// Note represents a personal note.
type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateNoteInput is the payload for creating a note.
type CreateNoteInput struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// UpdateNoteInput is the payload for updating a note (all fields optional).
type UpdateNoteInput struct {
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}
