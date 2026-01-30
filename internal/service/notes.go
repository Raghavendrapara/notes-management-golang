package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sort"
	"time"

	"go-backend/internal/models"
	"go-backend/internal/store"
)

var (
	ErrNoteNotFound = errors.New("note not found")
)

// NotesService contains business logic for notes.
type NotesService struct {
	store store.NoteStore
}

// NewNotesService returns a new notes service.
func NewNotesService(store store.NoteStore) *NotesService {
	return &NotesService{store: store}
}

// Create creates a new note with generated ID and timestamps.
func (s *NotesService) Create(in *models.CreateNoteInput) (*models.Note, error) {
	id, err := generateID()
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	note := &models.Note{
		ID:        id,
		Title:     in.Title,
		Body:      in.Body,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.store.Create(note); err != nil {
		return nil, err
	}
	return note, nil
}

// GetByID returns a note by ID.
func (s *NotesService) GetByID(id string) (*models.Note, error) {
	note, err := s.store.GetByID(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, ErrNoteNotFound
	}
	return note, nil
}

// List returns all notes, sorted by UpdatedAt descending.
func (s *NotesService) List() ([]*models.Note, error) {
	notes, err := s.store.List()
	if err != nil {
		return nil, err
	}
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].UpdatedAt.After(notes[j].UpdatedAt)
	})
	return notes, nil
}

// Update updates an existing note by ID.
func (s *NotesService) Update(id string, in *models.UpdateNoteInput) (*models.Note, error) {
	note, err := s.store.GetByID(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, ErrNoteNotFound
	}
	if in.Title != nil {
		note.Title = *in.Title
	}
	if in.Body != nil {
		note.Body = *in.Body
	}
	note.UpdatedAt = time.Now().UTC()
	if err := s.store.Update(note); err != nil {
		return nil, err
	}
	return note, nil
}

// Delete removes a note by ID.
func (s *NotesService) Delete(id string) error {
	exist, err := s.store.GetByID(id)
	if err != nil {
		return err
	}
	if exist == nil {
		return ErrNoteNotFound
	}
	return s.store.Delete(id)
}

func generateID() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
