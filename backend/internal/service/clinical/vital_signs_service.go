package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type VitalSignsService struct {
	repo *repository.VitalSignsRepository
}

func NewVitalSignsService(repo *repository.VitalSignsRepository) *VitalSignsService {
	return &VitalSignsService{repo: repo}
}

func (s *VitalSignsService) CreateVitalSigns(vitals *models.VitalSigns) (*models.VitalSigns, error) {
	// Auto-calculate BMI
	vitals.CalculateBMI()

	// Set measurement time if empty
	if vitals.MeasuredAt.IsZero() {
		vitals.MeasuredAt = time.Now()
	}

	return s.repo.Create(vitals)
}

func (s *VitalSignsService) GetVitalSignsByID(id uint) (*models.VitalSigns, error) {
	return s.repo.FindByID(id)
}

func (s *VitalSignsService) ListEncounterVitalSigns(encounterID uint) ([]*models.VitalSigns, error) {
	return s.repo.ListByEncounter(encounterID)
}

func (s *VitalSignsService) ListPatientVitalSigns(patientID uint, limit int) ([]*models.VitalSigns, error) {
	return s.repo.ListByPatient(patientID, limit)
}
