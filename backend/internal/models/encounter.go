package models

import (
	"time"
)

// Encounter represents an interaction between a patient and healthcare provider(s)
// FHIR R4 Encounter resource - Enhanced version
type Encounter struct {
	BaseModel

	// Status: planned, arrived, triaged, in-progress, onleave, finished, cancelled, entered-in-error
	Status string `gorm:"size:50;not null;index" json:"status"`

	// Class: imp (inpatient), amb (ambulatory), emer (emergency), hh (home health), vr (virtual)
	Class string `gorm:"size:50;not null" json:"class"`

	// Type: consultation, follow-up, emergency, etc.
	Type string `gorm:"size:100" json:"type,omitempty"`

	// Priority: routine, urgent, emergency
	Priority string `gorm:"size:50" json:"priority,omitempty"`

	// Patient
	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	// Practitioner (primary provider)
	PractitionerID *uint `gorm:"index" json:"practitioner_id,omitempty"`

	// Period
	PeriodStart time.Time  `gorm:"index;not null" json:"period_start"`
	PeriodEnd   *time.Time `json:"period_end,omitempty"`

	// Reason for visit
	Reason string `gorm:"type:text" json:"reason,omitempty"`

	// Diagnosis (can be multiple, but storing primary here)
	Diagnosis string `gorm:"type:text" json:"diagnosis,omitempty"`

	// Chief complaint
	ChiefComplaint string `gorm:"type:text" json:"chief_complaint,omitempty"`

	// Location
	LocationID *uint `gorm:"index" json:"location_id,omitempty"`

	// Service type: OPD, IPD, Outreach, Telemedicine
	ServiceType string `gorm:"size:100" json:"service_type,omitempty"`

	// Service Category: General, NCD Corner, MHPSS, MNCH & FP, Laboratory, Pharmacy, Emergency
	ServiceCategory string `gorm:"size:100" json:"service_category,omitempty"`

	// Appointment that created this encounter
	AppointmentID *uint        `gorm:"index" json:"appointment_id,omitempty"`
	Appointment   *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`

	// Relationships
	VitalSigns    []VitalSigns   `gorm:"foreignKey:EncounterID" json:"vital_signs,omitempty"`
	ClinicalNotes []ClinicalNote `gorm:"foreignKey:EncounterID" json:"clinical_notes,omitempty"`
	Prescriptions []Prescription `gorm:"foreignKey:EncounterID" json:"prescriptions,omitempty"`
	LabOrders     []LabOrder     `gorm:"foreignKey:EncounterID" json:"lab_orders,omitempty"`

	// Discharge information
	DischargeDisposition string     `gorm:"size:100" json:"discharge_disposition,omitempty"` // home, admitted, transferred, etc.
	DischargeDate        *time.Time `json:"discharge_date,omitempty"`
	DischargeSummary     string     `gorm:"type:text" json:"discharge_summary,omitempty"`

	// Follow-up
	FollowUpRequired bool       `gorm:"default:false" json:"follow_up_required"`
	FollowUpDate     *time.Time `json:"follow_up_date,omitempty"`
	FollowUpNotes    string     `gorm:"type:text" json:"follow_up_notes,omitempty"`
}

// TableName overrides the table name
func (Encounter) TableName() string {
	return "encounters"
}

// Start marks the encounter as in-progress
func (e *Encounter) Start() {
	e.Status = "in-progress"
}

// Finish completes the encounter
func (e *Encounter) Finish() {
	now := time.Now()
	e.Status = "finished"
	if e.PeriodEnd == nil {
		e.PeriodEnd = &now
	}
}

// Cancel cancels the encounter
func (e *Encounter) Cancel() {
	e.Status = "cancelled"
}

// GetDuration returns the encounter duration in minutes
func (e *Encounter) GetDuration() int {
	if e.PeriodEnd == nil {
		return 0
	}
	return int(e.PeriodEnd.Sub(e.PeriodStart).Minutes())
}

// IsActive checks if the encounter is currently active
func (e *Encounter) IsActive() bool {
	return e.Status == "in-progress" || e.Status == "arrived" || e.Status == "triaged"
}
