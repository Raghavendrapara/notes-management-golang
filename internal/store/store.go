package store

import (
	"sync"
	"time"

	"go-backend/internal/models"
)

// NoteStore defines the interface for note persistence.
// This allows swapping in-memory store for DB later without changing handlers/services.
type NoteStore interface {
	Create(note *models.Note) error
	GetByID(id string) (*models.Note, error)
	List() ([]*models.Note, error)
	Update(note *models.Note) error
	Delete(id string) error
}

// InMemoryNoteStore is an in-memory implementation for development/dummy data.
type InMemoryNoteStore struct {
	mu    sync.RWMutex
	notes map[string]*models.Note
}

// NewInMemoryNoteStore returns a new in-memory note store with optional seed data.
func NewInMemoryNoteStore() *InMemoryNoteStore {
	s := &InMemoryNoteStore{notes: make(map[string]*models.Note)}
	s.seed()
	return s
}

func (s *InMemoryNoteStore) seed() {
	now := time.Now().UTC()
	notes := []*models.Note{
		{ID: "1", Title: "Welcome", Body: "Your first note. Edit or add more via the API.", CreatedAt: now, UpdatedAt: now},
		{ID: "2", Title: "API Tips", Body: "GET /notes, POST /notes, PUT /notes/:id", CreatedAt: now, UpdatedAt: now},
	}
	for _, n := range notes {
		_ = s.Create(n)
	}
}

// Create saves a note. ID must be set by caller (or generate in service).
func (s *InMemoryNoteStore) Create(note *models.Note) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.notes[note.ID] = cloneNote(note)
	return nil
}

// GetByID returns a note by ID or nil if not found.
func (s *InMemoryNoteStore) GetByID(id string) (*models.Note, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n, ok := s.notes[id]
	if !ok {
		return nil, nil
	}
	return cloneNote(n), nil
}

// List returns all notes (order not guaranteed in map; sort in service if needed).
func (s *InMemoryNoteStore) List() ([]*models.Note, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]*models.Note, 0, len(s.notes))
	for _, n := range s.notes {
		out = append(out, cloneNote(n))
	}
	return out, nil
}

// Update overwrites an existing note.
func (s *InMemoryNoteStore) Update(note *models.Note) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.notes[note.ID]; !ok {
		return nil // consider ErrNotFound in production
	}
	s.notes[note.ID] = cloneNote(note)
	return nil
}

// Delete removes a note by ID.
func (s *InMemoryNoteStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.notes, id)
	return nil
}

func cloneNote(n *models.Note) *models.Note {
	c := *n
	return &c
}
