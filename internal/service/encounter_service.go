package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type EncounterService struct {
	repo *repository.EncounterRepository
}

func NewEncounterService(repo *repository.EncounterRepository) *EncounterService {
	return &EncounterService{repo: repo}
}

func (s *EncounterService) CreateEncounter(encounter *models.Encounter) (*models.Encounter, error) {
	// Set default status if empty
	if encounter.Status == "" {
		encounter.Status = "planned"
	}
	// Set default period start if empty
	if encounter.PeriodStart.IsZero() {
		encounter.PeriodStart = time.Now()
	}
	return s.repo.Create(encounter)
}

func (s *EncounterService) GetEncounterByID(id uint) (*models.Encounter, error) {
	return s.repo.FindByID(id)
}

func (s *EncounterService) UpdateEncounter(encounter *models.Encounter) (*models.Encounter, error) {
	return s.repo.Update(encounter)
}

func (s *EncounterService) ListPatientEncounters(patientID uint, page, limit int) ([]*models.Encounter, int64, error) {
	offset := (page - 1) * limit
	return s.repo.ListByPatient(patientID, offset, limit)
}

func (s *EncounterService) StartEncounter(id uint) (*models.Encounter, error) {
	encounter, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	encounter.Start()
	return s.repo.Update(encounter)
}

func (s *EncounterService) FinishEncounter(id uint) (*models.Encounter, error) {
	encounter, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	encounter.Finish()
	return s.repo.Update(encounter)
}
