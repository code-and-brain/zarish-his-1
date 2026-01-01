package models

import (
	"time"
)

// LabTest represents a laboratory test in the catalog
type LabTest struct {
	BaseModel

	Code     string `gorm:"size:50;uniqueIndex;not null" json:"code"` // LOINC code or local code
	Name     string `gorm:"size:255;not null;index" json:"name"`
	Category string `gorm:"size:100;index" json:"category"` // hematology, biochemistry, microbiology, etc.

	// Unit of measurement
	Unit string `gorm:"size:50" json:"unit,omitempty"`

	// Reference ranges
	ReferenceRangeMin *float64 `json:"reference_range_min,omitempty"`
	ReferenceRangeMax *float64 `json:"reference_range_max,omitempty"`
	ReferenceText     string   `gorm:"type:text" json:"reference_text,omitempty"` // For non-numeric ranges

	// Sample type: blood, urine, stool, etc.
	SampleType string `gorm:"size:100" json:"sample_type,omitempty"`

	// Turnaround time in hours
	TurnaroundTime int `json:"turnaround_time,omitempty"`

	// Cost
	Cost float64 `json:"cost,omitempty"`

	Active bool   `gorm:"default:true" json:"active"`
	Notes  string `gorm:"type:text" json:"notes,omitempty"`
}

// TableName overrides the table name
func (LabTest) TableName() string {
	return "lab_tests"
}

// LabOrder represents a laboratory test order
// FHIR R4 ServiceRequest resource
type LabOrder struct {
	BaseModel

	EncounterID uint      `gorm:"index;not null" json:"encounter_id"`
	Encounter   Encounter `gorm:"foreignKey:EncounterID" json:"encounter,omitempty"`

	PatientID uint    `gorm:"index;not null" json:"patient_id"`
	Patient   Patient `gorm:"foreignKey:PatientID" json:"patient,omitempty"`

	PractitionerID uint `gorm:"index" json:"practitioner_id,omitempty"` // Ordering physician

	OrderDate time.Time `gorm:"not null;index" json:"order_date"`

	// Status: ordered, collected, processing, completed, cancelled
	Status string `gorm:"size:50;not null;default:'ordered';index" json:"status"`

	// Priority: routine, urgent, stat
	Priority string `gorm:"size:50;default:'routine'" json:"priority"`

	// Clinical information
	ClinicalInfo string `gorm:"type:text" json:"clinical_info,omitempty"`

	// Sample collection
	SampleCollectedAt *time.Time `json:"sample_collected_at,omitempty"`
	SampleCollectedBy *uint      `json:"sample_collected_by,omitempty"`

	// Results
	ResultsAvailableAt *time.Time `json:"results_available_at,omitempty"`
	ResultsReviewedBy  *uint      `json:"results_reviewed_by,omitempty"`
	ResultsReviewedAt  *time.Time `json:"results_reviewed_at,omitempty"`

	// Relationships
	Results []LabResult `gorm:"foreignKey:LabOrderID" json:"results,omitempty"`

	Notes string `gorm:"type:text" json:"notes,omitempty"`
}

// TableName overrides the table name
func (LabOrder) TableName() string {
	return "lab_orders"
}

// LabResult represents a single test result within a lab order
// FHIR R4 Observation resource (category: laboratory)
type LabResult struct {
	BaseModel

	LabOrderID uint     `gorm:"index;not null" json:"lab_order_id"`
	LabOrder   LabOrder `gorm:"foreignKey:LabOrderID" json:"lab_order,omitempty"`

	LabTestID uint    `gorm:"index;not null" json:"lab_test_id"`
	LabTest   LabTest `gorm:"foreignKey:LabTestID" json:"lab_test,omitempty"`

	// Result value (can be numeric or text)
	Value        string   `gorm:"size:255" json:"value"`
	NumericValue *float64 `json:"numeric_value,omitempty"`
	Unit         string   `gorm:"size:50" json:"unit,omitempty"`

	// Abnormal flag: normal, high, low, critical-high, critical-low
	AbnormalFlag string `gorm:"size:20" json:"abnormal_flag,omitempty"`

	// Reference range for this specific result
	ReferenceRange string `gorm:"size:255" json:"reference_range,omitempty"`

	// Result date/time
	ResultDate time.Time `gorm:"not null;index" json:"result_date"`

	// Interpretation/comments
	Interpretation string `gorm:"type:text" json:"interpretation,omitempty"`
	Notes          string `gorm:"type:text" json:"notes,omitempty"`

	// Performed by
	PerformedBy uint `json:"performed_by,omitempty"`

	// Status: preliminary, final, corrected, cancelled
	Status string `gorm:"size:50;default:'final'" json:"status"`
}

// TableName overrides the table name
func (LabResult) TableName() string {
	return "lab_results"
}

// DetermineAbnormalFlag automatically determines if the result is abnormal
func (lr *LabResult) DetermineAbnormalFlag(test *LabTest) {
	if lr.NumericValue == nil || test == nil {
		lr.AbnormalFlag = ""
		return
	}

	value := *lr.NumericValue

	// Check critical ranges first (if defined separately)
	// For now, using reference ranges

	if test.ReferenceRangeMin != nil && value < *test.ReferenceRangeMin {
		lr.AbnormalFlag = "low"
		// Could add critical-low threshold check here
		return
	}

	if test.ReferenceRangeMax != nil && value > *test.ReferenceRangeMax {
		lr.AbnormalFlag = "high"
		// Could add critical-high threshold check here
		return
	}

	lr.AbnormalFlag = "normal"
}
