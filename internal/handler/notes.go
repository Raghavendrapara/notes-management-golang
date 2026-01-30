package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go-backend/internal/models"
	"go-backend/internal/service"
)

// NotesHandler handles HTTP for notes.
type NotesHandler struct {
	svc *service.NotesService
}

// NewNotesHandler returns a new notes HTTP handler.
func NewNotesHandler(svc *service.NotesService) *NotesHandler {
	return &NotesHandler{svc: svc}
}

// List handles GET /notes.
func (h *NotesHandler) List(w http.ResponseWriter, r *http.Request) {
	notes, err := h.svc.List()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if notes == nil {
		notes = []*models.Note{}
	}
	respondJSON(w, http.StatusOK, notes)
}

// GetByID handles GET /notes/{id}.
func (h *NotesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "missing id")
		return
	}
	note, err := h.svc.GetByID(id)
	if err != nil {
		if err == service.ErrNoteNotFound {
			respondError(w, http.StatusNotFound, "note not found")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, note)
}

// Create handles POST /notes.
func (h *NotesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var in models.CreateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if in.Title == "" {
		respondError(w, http.StatusBadRequest, "title is required")
		return
	}
	note, err := h.svc.Create(&in)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, note)
}

// Update handles PUT /notes/{id}.
func (h *NotesHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "missing id")
		return
	}
	var in models.UpdateNoteInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	note, err := h.svc.Update(id, &in)
	if err != nil {
		if err == service.ErrNoteNotFound {
			respondError(w, http.StatusNotFound, "note not found")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, note)
}

// Delete handles DELETE /notes/{id}.
func (h *NotesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "missing id")
		return
	}
	if err := h.svc.Delete(id); err != nil {
		if err == service.ErrNoteNotFound {
			respondError(w, http.StatusNotFound, "note not found")
			return
		}
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
