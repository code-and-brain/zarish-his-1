package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type MedicationService struct {
	repo       *repository.MedicationRepository
	cdsService *CDSService
}

func NewMedicationService(repo *repository.MedicationRepository, cdsService *CDSService) *MedicationService {
	return &MedicationService{
		repo:       repo,
		cdsService: cdsService,
	}
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

// CheckPrescriptionSafety validates a prescription against CDS rules
func (s *MedicationService) CheckPrescriptionSafety(prescription *models.Prescription, patient *models.Patient) ([]string, error) {
	var warnings []string

	// 1. Check Allergies
	// Need to fetch medication name first if not present
	med, err := s.repo.FindMedicationByID(prescription.MedicationID)
	if err != nil {
		return nil, err
	}

	allergyWarnings := s.cdsService.CheckAllergies(med.Name, patient.Allergies)
	warnings = append(warnings, allergyWarnings...)

	// 2. Check Interactions
	activeRx, err := s.repo.ListActivePrescriptions(prescription.PatientID)
	if err != nil {
		return nil, err
	}

	// Convert []*models.Prescription to []models.Prescription for the service
	var activeRxList []models.Prescription
	for _, rx := range activeRx {
		activeRxList = append(activeRxList, *rx)
	}

	interactionWarnings := s.cdsService.CheckInteractions(prescription.MedicationID, activeRxList)
	warnings = append(warnings, interactionWarnings...)

	return warnings, nil
}
