package service

import (
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	// Add logic to generate MRN if missing
	if patient.MRN == "" {
		// Generate MRN logic here
	}
	return s.repo.CreatePatient(patient)
}

func (s *PatientService) GetPatient(id uint) (*models.Patient, error) {
	return s.repo.GetPatientByID(id)
}

func (s *PatientService) ListPatients(page, pageSize int) ([]models.Patient, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListPatients(offset, pageSize)
}
