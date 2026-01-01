package models

import (
	"time"
)

// ClinicalNote represents clinical documentation for an encounter
// Supports SOAP (Subjective, Objective, Assessment, Plan) format
type ClinicalNote struct {
	BaseModel

	EncounterID uint      `gorm:"index;not null" json:"encounter_id"`
	Encounter   Encounter `gorm:"foreignKey:EncounterID" json:"encounter,omitempty"`

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	PractitionerID uint `gorm:"index" json:"practitioner_id,omitempty"` // Doctor/Provider ID

	NoteType string `gorm:"size:50;not null" json:"note_type"` // soap, progress, discharge, admission, consultation

	// SOAP Components
	Subjective string `gorm:"type:text" json:"subjective,omitempty"` // Patient's complaints and symptoms
	Objective  string `gorm:"type:text" json:"objective,omitempty"`  // Physical examination findings
	Assessment string `gorm:"type:text" json:"assessment,omitempty"` // Diagnosis and clinical impression
	Plan       string `gorm:"type:text" json:"plan,omitempty"`       // Treatment plan and follow-up

	// Additional fields
	ChiefComplaint        string `gorm:"type:text" json:"chief_complaint,omitempty"`
	HistoryPresentIllness string `gorm:"type:text" json:"history_present_illness,omitempty"`
	PhysicalExamination   string `gorm:"type:text" json:"physical_examination,omitempty"`
	ReviewOfSystems       string `gorm:"type:text" json:"review_of_systems,omitempty"`
	DifferentialDiagnosis string `gorm:"type:text" json:"differential_diagnosis,omitempty"`
	TreatmentPlan         string `gorm:"type:text" json:"treatment_plan,omitempty"`
	FollowUpInstructions  string `gorm:"type:text" json:"follow_up_instructions,omitempty"`

	// Metadata
	NoteDate time.Time `gorm:"not null;index" json:"note_date"`
	Status   string    `gorm:"size:50;default:'draft'" json:"status"` // draft, final, amended

	// Signature
	SignedBy  *uint      `json:"signed_by,omitempty"` // User ID who signed
	SignedAt  *time.Time `json:"signed_at,omitempty"`
	IsAmended bool       `gorm:"default:false" json:"is_amended"`
	AmendedAt *time.Time `json:"amended_at,omitempty"`
	AmendedBy *uint      `json:"amended_by,omitempty"`
}

// TableName overrides the table name
func (ClinicalNote) TableName() string {
	return "clinical_notes"
}

// Sign marks the note as final and signed
func (c *ClinicalNote) Sign(userID uint) {
	now := time.Now()
	c.Status = "final"
	c.SignedBy = &userID
	c.SignedAt = &now
}

// Amend marks the note as amended
func (c *ClinicalNote) Amend(userID uint) {
	now := time.Now()
	c.IsAmended = true
	c.AmendedBy = &userID
	c.AmendedAt = &now
}
