package service

import (
	"time"

	"github.com/ZarishSphere-Platform/zarish-his/internal/models"
	"github.com/ZarishSphere-Platform/zarish-his/internal/repository"
)

type ClinicalNoteService struct {
	repo *repository.ClinicalNoteRepository
}

func NewClinicalNoteService(repo *repository.ClinicalNoteRepository) *ClinicalNoteService {
	return &ClinicalNoteService{repo: repo}
}

func (s *ClinicalNoteService) CreateNote(note *models.ClinicalNote) (*models.ClinicalNote, error) {
	if note.NoteDate.IsZero() {
		note.NoteDate = time.Now()
	}
	if note.Status == "" {
		note.Status = "draft"
	}
	return s.repo.Create(note)
}

func (s *ClinicalNoteService) GetNoteByID(id uint) (*models.ClinicalNote, error) {
	return s.repo.FindByID(id)
}

func (s *ClinicalNoteService) UpdateNote(note *models.ClinicalNote) (*models.ClinicalNote, error) {
	return s.repo.Update(note)
}

func (s *ClinicalNoteService) SignNote(id uint, userID uint) (*models.ClinicalNote, error) {
	note, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	note.Sign(userID)
	return s.repo.Update(note)
}

func (s *ClinicalNoteService) ListEncounterNotes(encounterID uint) ([]*models.ClinicalNote, error) {
	return s.repo.ListByEncounter(encounterID)
}

func (s *ClinicalNoteService) ListPatientNotes(patientID uint, limit int) ([]*models.ClinicalNote, error) {
	return s.repo.ListByPatient(patientID, limit)
}
