package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type MedicationService struct {
	repo *repository.MedicationRepository
}

func NewMedicationService(repo *repository.MedicationRepository) *MedicationService {
	return &MedicationService{repo: repo}
}

// Medication methods
func (s *MedicationService) CreateMedication(med *models.Medication) (*models.Medication, error) {
	return s.repo.CreateMedication(med)
}

func (s *MedicationService) SearchMedications(query string) ([]*models.Medication, error) {
	return s.repo.SearchMedications(query)
}

// Prescription methods
func (s *MedicationService) CreatePrescription(prescription *models.Prescription) (*models.Prescription, error) {
	if prescription.StartDate.IsZero() {
		prescription.StartDate = time.Now()
	}
	if prescription.Status == "" {
		prescription.Status = "active"
	}
	return s.repo.CreatePrescription(prescription)
}

func (s *MedicationService) GetPrescriptionByID(id uint) (*models.Prescription, error) {
	return s.repo.FindPrescriptionByID(id)
}

func (s *MedicationService) DiscontinuePrescription(id uint, reason string) (*models.Prescription, error) {
	prescription, err := s.repo.FindPrescriptionByID(id)
	if err != nil {
		return nil, err
	}
	prescription.Discontinue(reason)
	return s.repo.UpdatePrescription(prescription)
}

func (s *MedicationService) ListPatientPrescriptions(patientID uint) ([]*models.Prescription, error) {
	return s.repo.ListPrescriptionsByPatient(patientID)
}

func (s *MedicationService) ListActivePrescriptions(patientID uint) ([]*models.Prescription, error) {
	return s.repo.ListActivePrescriptions(patientID)
}
