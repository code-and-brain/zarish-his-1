package service

import (
	"errors"
	"fmt"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type PatientService struct {
	repo *repository.PatientRepository
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

func (s *PatientService) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	// Generate MRN if not provided
	if patient.MRN == "" {
		mrn, err := s.generateMRN()
		if err != nil {
			return nil, err
		}
		patient.MRN = mrn
	}

	return s.repo.Create(patient)
}

func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
	return s.repo.FindByID(id)
}

func (s *PatientService) UpdatePatient(patient *models.Patient) (*models.Patient, error) {
	return s.repo.Update(patient)
}

func (s *PatientService) DeletePatient(id uint) error {
	return s.repo.Delete(id)
}

func (s *PatientService) ListPatients(page, limit int, nationality, search string) ([]*models.Patient, int64, error) {
	offset := (page - 1) * limit
	return s.repo.List(offset, limit, nationality, search)
}

func (s *PatientService) SearchPatients(query string) ([]*models.Patient, error) {
	return s.repo.Search(query)
}

func (s *PatientService) GetPatientHistory(id uint) (map[string]interface{}, error) {
	patient, err := s.repo.FindByIDWithRelations(id)
	if err != nil {
		return nil, err
	}

	history := map[string]interface{}{
		"patient":       patient,
		"encounters":    patient.Encounters,
		"appointments":  patient.Appointments,
		"prescriptions": patient.Prescriptions,
		"lab_orders":    patient.LabOrders,
	}

	return history, nil
}

// generateMRN generates a unique Medical Record Number
func (s *PatientService) generateMRN() (string, error) {
	// Get the last patient to determine the next MRN
	lastPatient, err := s.repo.GetLastPatient()
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return "", err
	}

	var nextNumber int
	if lastPatient != nil && lastPatient.MRN != "" {
		// Extract number from MRN (assuming format: MRN-XXXXXX)
		fmt.Sscanf(lastPatient.MRN, "MRN-%d", &nextNumber)
		nextNumber++
	} else {
		nextNumber = 1
	}

	return fmt.Sprintf("MRN-%06d", nextNumber), nil
}
