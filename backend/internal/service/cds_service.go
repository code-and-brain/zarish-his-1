package service

import (
	"strings"

	"github.com/zarishsphere/zarish-his/internal/models"
)

type CDSService struct{}

func NewCDSService() *CDSService {
	return &CDSService{}
}

// CheckInteractions checks for drug-drug interactions
func (s *CDSService) CheckInteractions(newMedID uint, activePrescriptions []models.Prescription) []string {
	var warnings []string

	// Mock interaction database
	// In a real system, this would query a knowledge base
	interactions := map[uint][]uint{
		1: {2, 3}, // Med 1 interacts with 2 and 3
		2: {1},    // Med 2 interacts with 1
	}

	if interactingMeds, exists := interactions[newMedID]; exists {
		for _, activeRx := range activePrescriptions {
			for _, interactingID := range interactingMeds {
				if activeRx.MedicationID == interactingID {
					warnings = append(warnings, "Potential interaction with active prescription: "+activeRx.Medication.Name)
				}
			}
		}
	}

	return warnings
}

// CheckAllergies checks if the patient is allergic to the medication
func (s *CDSService) CheckAllergies(medName string, allergies []string) []string {
	var warnings []string
	medNameLower := strings.ToLower(medName)

	for _, allergy := range allergies {
		if strings.Contains(medNameLower, strings.ToLower(allergy)) {
			warnings = append(warnings, "Patient has reported allergy to: "+allergy)
		}
	}

	return warnings
}
