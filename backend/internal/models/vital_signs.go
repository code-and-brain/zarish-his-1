package models

import (
	"time"
)

// VitalSigns represents vital signs measurements linked to an encounter
// FHIR R4 Observation resource (category: vital-signs)
type VitalSigns struct {
	BaseModel

	EncounterID uint      `gorm:"index;not null" json:"encounter_id"`
	Encounter   Encounter `gorm:"foreignKey:EncounterID" json:"encounter,omitempty"`

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	MeasuredAt time.Time `gorm:"not null;index" json:"measured_at"`

	// Blood Pressure (mmHg)
	SystolicBP  *int `json:"systolic_bp,omitempty"`  // Normal: 90-120
	DiastolicBP *int `json:"diastolic_bp,omitempty"` // Normal: 60-80

	// Pulse/Heart Rate (bpm)
	PulseRate *int `json:"pulse_rate,omitempty"` // Normal: 60-100

	// Respiratory Rate (breaths/min)
	RespiratoryRate *int `json:"respiratory_rate,omitempty"` // Normal: 12-20

	// Temperature (°C)
	Temperature *float64 `json:"temperature,omitempty"` // Normal: 36.1-37.2

	// Oxygen Saturation (%)
	SpO2 *int `json:"spo2,omitempty"` // Normal: 95-100

	// Weight (kg)
	Weight *float64 `json:"weight,omitempty"`

	// Height (cm)
	Height *float64 `json:"height,omitempty"`

	// BMI (calculated)
	BMI *float64 `json:"bmi,omitempty"`

	// Pain Scale (0-10)
	PainScale *int `json:"pain_scale,omitempty"`

	// Notes
	Notes string `gorm:"type:text" json:"notes,omitempty"`

	// Recorded by
	RecordedBy uint `json:"recorded_by,omitempty"` // User ID of the person who recorded
}

// TableName overrides the table name
func (VitalSigns) TableName() string {
	return "vital_signs"
}

// CalculateBMI calculates BMI from weight and height
func (v *VitalSigns) CalculateBMI() {
	if v.Weight != nil && v.Height != nil && *v.Height > 0 {
		heightInMeters := *v.Height / 100.0
		bmi := *v.Weight / (heightInMeters * heightInMeters)
		v.BMI = &bmi
	}
}

// IsAbnormal checks if any vital signs are outside normal ranges
func (v *VitalSigns) IsAbnormal() bool {
	if v.SystolicBP != nil && (*v.SystolicBP < 90 || *v.SystolicBP > 140) {
		return true
	}
	if v.DiastolicBP != nil && (*v.DiastolicBP < 60 || *v.DiastolicBP > 90) {
		return true
	}
	if v.PulseRate != nil && (*v.PulseRate < 60 || *v.PulseRate > 100) {
		return true
	}
	if v.Temperature != nil && (*v.Temperature < 36.1 || *v.Temperature > 37.8) {
		return true
	}
	if v.SpO2 != nil && *v.SpO2 < 95 {
		return true
	}
	return false
}
