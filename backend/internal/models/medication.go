package models

import (
	"time"
)

// Medication represents a medication in the formulary
type Medication struct {
	BaseModel

	Name        string `gorm:"size:255;not null;index" json:"name"`
	GenericName string `gorm:"size:255;index" json:"generic_name,omitempty"`
	BrandName   string `gorm:"size:255" json:"brand_name,omitempty"`

	// Form: tablet, capsule, syrup, injection, cream, drops, etc.
	Form string `gorm:"size:100" json:"form"`

	// Strength: e.g., "500mg", "10mg/ml"
	Strength string `gorm:"size:100" json:"strength"`

	// Unit: mg, ml, IU, etc.
	Unit string `gorm:"size:50" json:"unit"`

	// Category: antibiotic, analgesic, antihypertensive, etc.
	Category string `gorm:"size:100;index" json:"category,omitempty"`

	// Active ingredient
	ActiveIngredient string `gorm:"size:255" json:"active_ingredient,omitempty"`

	// Manufacturer
	Manufacturer string `gorm:"size:255" json:"manufacturer,omitempty"`

	// Status
	Active bool `gorm:"default:true" json:"active"`

	// Notes
	Notes string `gorm:"type:text" json:"notes,omitempty"`
}

// TableName overrides the table name
func (Medication) TableName() string {
	return "medications"
}

// Prescription represents a medication prescription
// FHIR R4 MedicationRequest resource
type Prescription struct {
	BaseModel

	EncounterID uint      `gorm:"index;not null" json:"encounter_id"`
	Encounter   Encounter `gorm:"foreignKey:EncounterID" json:"encounter,omitempty"`

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	MedicationID uint       `gorm:"index;not null" json:"medication_id"`
	Medication   Medication `gorm:"foreignKey:MedicationID" json:"medication,omitempty"`

	PractitionerID uint `gorm:"index" json:"practitioner_id,omitempty"` // Prescriber ID

	// Dosage: e.g., "1 tablet", "5ml", "2 puffs"
	Dosage string `gorm:"size:100;not null" json:"dosage"`

	// Frequency: e.g., "twice daily", "every 6 hours", "as needed"
	Frequency string `gorm:"size:100;not null" json:"frequency"`

	// Route: oral, IV, IM, topical, inhalation, etc.
	Route string `gorm:"size:50" json:"route"`

	// Duration in days
	DurationDays int `json:"duration_days"`

	// Quantity prescribed
	Quantity int `json:"quantity"`

	// Number of refills allowed
	Refills int `gorm:"default:0" json:"refills"`

	// Instructions for patient
	Instructions string `gorm:"type:text" json:"instructions,omitempty"`

	// Special instructions (e.g., "take with food", "avoid alcohol")
	SpecialInstructions string `gorm:"type:text" json:"special_instructions,omitempty"`

	// Dates
	StartDate time.Time  `gorm:"not null;index" json:"start_date"`
	EndDate   *time.Time `json:"end_date,omitempty"`

	// Status: active, completed, discontinued, cancelled
	Status string `gorm:"size:50;not null;default:'active';index" json:"status"`

	// Reason for discontinuation
	DiscontinuedReason string     `gorm:"type:text" json:"discontinued_reason,omitempty"`
	DiscontinuedAt     *time.Time `json:"discontinued_at,omitempty"`

	// Dispensed information
	DispensedDate *time.Time `json:"dispensed_date,omitempty"`
	DispensedBy   *uint      `json:"dispensed_by,omitempty"`
}

// TableName overrides the table name
func (Prescription) TableName() string {
	return "prescriptions"
}

// Discontinue marks the prescription as discontinued
func (p *Prescription) Discontinue(reason string) {
	now := time.Now()
	p.Status = "discontinued"
	p.DiscontinuedReason = reason
	p.DiscontinuedAt = &now
}

// Complete marks the prescription as completed
func (p *Prescription) Complete() {
	p.Status = "completed"
	if p.EndDate == nil {
		now := time.Now()
		p.EndDate = &now
	}
}

// IsActive checks if the prescription is currently active
func (p *Prescription) IsActive() bool {
	if p.Status != "active" {
		return false
	}
	if p.EndDate != nil && p.EndDate.Before(time.Now()) {
		return false
	}
	return true
}
