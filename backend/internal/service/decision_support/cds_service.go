package decision_support

import "zarish-his/backend/internal/domain/models"

// CDS stands for Clinical Decision Support

type CDSService struct {
	// Depending on the complexity, this might need access to various repositories
}

func NewCDSService() *CDSService {
	return &CDSService{}
}

func (s *CDSService) CheckDrugInteractions(medications []models.Medication) ([]string, error) {
	// In a real implementation, this would check against a drug interaction database.
	warnings := []string{}
	if len(medications) > 1 {
		warnings = append(warnings, "Potential drug-drug interaction. Please review patient's medications.")
	}
	return warnings, nil
}

func (s *CDSService) CheckAllergies(patient models.Patient, medication models.Medication) ([]string, error) {
	// In a real implementation, this would check against a patient's allergy list.
	warnings := []string{}
	// Dummy check
	if patient.ID == 1 && medication.Name == "Penicillin" {
		warnings = append(warnings, "Patient has a known allergy to Penicillin.")
	}
	return warnings, nil
}
