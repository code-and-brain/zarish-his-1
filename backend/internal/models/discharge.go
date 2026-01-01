package models

import "time"

// DischargeSummary represents a patient discharge summary
type DischargeSummary struct {
	BaseModel

	AdmissionID uint       `gorm:"not null;index" json:"admission_id"`
	Admission   *Admission `gorm:"foreignKey:AdmissionID" json:"admission,omitempty"`

	DischargeDate time.Time `gorm:"not null" json:"discharge_date"`
	DischargeType string    `gorm:"size:50;not null" json:"discharge_type"` // Regular, AMA, Transfer, Death

	ChiefComplaint         string `gorm:"type:text" json:"chief_complaint"`
	Diagnosis              string `gorm:"type:text;not null" json:"diagnosis"`
	TreatmentSummary       string `gorm:"type:text" json:"treatment_summary"`
	MedicationsOnDischarge string `gorm:"type:text" json:"medications_on_discharge"`
	FollowUpInstructions   string `gorm:"type:text" json:"follow_up_instructions"`

	SignedBy uint `gorm:"not null" json:"signed_by"` // UserID
	// User     *User `gorm:"foreignKey:SignedBy" json:"signed_by_user,omitempty"`
}

func (DischargeSummary) TableName() string {
	return "discharge_summaries"
}
