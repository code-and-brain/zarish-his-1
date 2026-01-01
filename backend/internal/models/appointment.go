package models

import (
	"time"
)

// Appointment represents a scheduled appointment
// FHIR R4 Appointment resource
type Appointment struct {
	BaseModel

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	PractitionerID *uint `gorm:"index" json:"practitioner_id,omitempty"` // Doctor/Provider

	// Appointment type: consultation, follow-up, procedure, vaccination, etc.
	AppointmentType string `gorm:"size:100" json:"appointment_type"`

	// Status: scheduled, confirmed, arrived, in-progress, completed, cancelled, no-show
	Status string `gorm:"size:50;not null;default:'scheduled';index" json:"status"`

	// Scheduled times
	ScheduledStart time.Time `gorm:"not null;index" json:"scheduled_start"`
	ScheduledEnd   time.Time `gorm:"not null;index" json:"scheduled_end"`

	// Actual times
	ActualStart *time.Time `json:"actual_start,omitempty"`
	ActualEnd   *time.Time `json:"actual_end,omitempty"`

	// Reason for visit
	Reason string `gorm:"type:text" json:"reason,omitempty"`

	// Notes
	Notes string `gorm:"type:text" json:"notes,omitempty"`

	// Cancellation
	CancelledAt        *time.Time `json:"cancelled_at,omitempty"`
	CancelledBy        *uint      `json:"cancelled_by,omitempty"`
	CancellationReason string     `gorm:"type:text" json:"cancellation_reason,omitempty"`

	// Reminder sent
	ReminderSent   bool       `gorm:"default:false" json:"reminder_sent"`
	ReminderSentAt *time.Time `json:"reminder_sent_at,omitempty"`

	// Created by (receptionist/admin)
	CreatedBy uint `json:"created_by,omitempty"`

	// Link to encounter (if appointment resulted in an encounter)
	EncounterID *uint      `gorm:"index" json:"encounter_id,omitempty"`
	Encounter   *Encounter `gorm:"foreignKey:EncounterID" json:"encounter,omitempty"`
}

// TableName overrides the table name
func (Appointment) TableName() string {
	return "appointments"
}

// Cancel cancels the appointment
func (a *Appointment) Cancel(reason string, cancelledBy uint) {
	now := time.Now()
	a.Status = "cancelled"
	a.CancellationReason = reason
	a.CancelledAt = &now
	a.CancelledBy = &cancelledBy
}

// MarkArrived marks the patient as arrived
func (a *Appointment) MarkArrived() {
	a.Status = "arrived"
}

// Start starts the appointment
func (a *Appointment) Start() {
	now := time.Now()
	a.Status = "in-progress"
	a.ActualStart = &now
}

// Complete completes the appointment
func (a *Appointment) Complete() {
	now := time.Now()
	a.Status = "completed"
	if a.ActualStart == nil {
		a.ActualStart = &now
	}
	a.ActualEnd = &now
}

// MarkNoShow marks the patient as no-show
func (a *Appointment) MarkNoShow() {
	a.Status = "no-show"
}

// GetDuration returns the scheduled duration in minutes
func (a *Appointment) GetDuration() int {
	return int(a.ScheduledEnd.Sub(a.ScheduledStart).Minutes())
}

// IsUpcoming checks if the appointment is in the future
func (a *Appointment) IsUpcoming() bool {
	return a.ScheduledStart.After(time.Now()) && (a.Status == "scheduled" || a.Status == "confirmed")
}
