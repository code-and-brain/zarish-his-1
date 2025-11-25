package models

import (
	"time"
)

// Encounter represents an interaction between a patient and healthcare provider(s).
// Simplified from FHIR Encounter resource.
type Encounter struct {
	BaseModel

	Status string `gorm:"size:50;not null;index" json:"status"` // planned, arrived, triaged, in-progress, onleave, finished, cancelled
	Class  string `gorm:"size:50;not null" json:"class"`        // imp (inpatient), amb (ambulatory), emer (emergency)
	Type   string `gorm:"size:100" json:"type"`                 // Specific type of encounter

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	PractitionerID uint `gorm:"index" json:"practitioner_id"` // Primary provider

	PeriodStart time.Time  `gorm:"index;not null" json:"period_start"`
	PeriodEnd   *time.Time `json:"period_end,omitempty"`

	Reason    string `gorm:"type:text" json:"reason"`
	Diagnosis string `gorm:"type:text" json:"diagnosis"`

	LocationID *uint `gorm:"index" json:"location_id,omitempty"`
}

// TableName overrides the table name used by User to `encounters`
func (Encounter) TableName() string {
	return "encounters"
}
