package models

import (
	"time"
)

// Patient represents a patient record.
// Simplified from FHIR Patient resource.
type Patient struct {
	BaseModel

	Active bool   `gorm:"default:true" json:"active"`
	MRN    string `gorm:"size:64;uniqueIndex;not null" json:"mrn"` // Medical Record Number

	FirstName string    `gorm:"size:100" json:"first_name"`
	LastName  string    `gorm:"size:100" json:"last_name"`
	Gender    string    `gorm:"size:20" json:"gender"`
	BirthDate time.Time `json:"birth_date"`

	Phone string `gorm:"size:50" json:"phone"`
	Email string `gorm:"size:255" json:"email"`

	Address    string `gorm:"type:text" json:"address"`
	City       string `gorm:"size:100" json:"city"`
	State      string `gorm:"size:100" json:"state"`
	PostalCode string `gorm:"size:20" json:"postal_code"`
	Country    string `gorm:"size:100" json:"country"`

	MaritalStatus string `gorm:"size:50" json:"marital_status"`
}

// TableName overrides the table name used by User to `patients`
func (Patient) TableName() string {
	return "patients"
}
