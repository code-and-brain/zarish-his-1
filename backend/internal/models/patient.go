package models

import (
	"time"
)

// Patient represents a FHIR R4-compliant patient record with Bangladesh-specific extensions
type Patient struct {
	BaseModel

	// Core FHIR fields
	Active bool   `gorm:"default:true" json:"active"`
	MRN    string `gorm:"size:64;uniqueIndex;not null" json:"mrn"` // Medical Record Number

	// Name (FHIR HumanName)
	GivenName  string `gorm:"size:100" json:"given_name"`  // First name
	FamilyName string `gorm:"size:100" json:"family_name"` // Last name
	MiddleName string `gorm:"size:100" json:"middle_name,omitempty"`

	// Demographics
	Gender    string     `gorm:"size:20" json:"gender"` // male, female, other, unknown
	BirthDate *time.Time `json:"birth_date"`

	// Bangladesh-specific: Nationality
	Nationality string `gorm:"size:50;index" json:"nationality"` // bangladeshi, rohingya, other

	// Conditional fields based on nationality
	// For Bangladeshi citizens
	NationalID string `gorm:"size:50;index" json:"national_id,omitempty"` // NID number
	BirthRegNo string `gorm:"size:50" json:"birth_reg_no,omitempty"`      // Birth registration number

	// For Rohingya refugees
	UNHCRNumber string `gorm:"size:50;index" json:"unhcr_number,omitempty"` // UNHCR registration number
	CampName    string `gorm:"size:100" json:"camp_name,omitempty"`         // Refugee camp name
	BlockNumber string `gorm:"size:50" json:"block_number,omitempty"`       // Camp block number

	// Contact information (FHIR ContactPoint)
	Phone  string `gorm:"size:50" json:"phone"`
	Email  string `gorm:"size:255" json:"email,omitempty"`
	Phone2 string `gorm:"size:50" json:"phone2,omitempty"` // Alternative phone

	// Address (FHIR Address)
	AddressLine1 string `gorm:"type:text" json:"address_line1,omitempty"`
	AddressLine2 string `gorm:"type:text" json:"address_line2,omitempty"`
	City         string `gorm:"size:100" json:"city,omitempty"`
	District     string `gorm:"size:100" json:"district,omitempty"` // Bangladesh administrative division
	Division     string `gorm:"size:100" json:"division,omitempty"` // Bangladesh division
	PostalCode   string `gorm:"size:20" json:"postal_code,omitempty"`
	Country      string `gorm:"size:100;default:'Bangladesh'" json:"country"`

	// Additional demographics
	MaritalStatus string `gorm:"size:50" json:"marital_status,omitempty"` // single, married, divorced, widowed
	Occupation    string `gorm:"size:100" json:"occupation,omitempty"`
	Religion      string `gorm:"size:50" json:"religion,omitempty"` // islam, hinduism, buddhism, christianity, other

	// Emergency contact
	EmergencyContactName     string `gorm:"size:200" json:"emergency_contact_name,omitempty"`
	EmergencyContactPhone    string `gorm:"size:50" json:"emergency_contact_phone,omitempty"`
	EmergencyContactRelation string `gorm:"size:50" json:"emergency_contact_relation,omitempty"`

	// Photo
	PhotoURL string `gorm:"size:500" json:"photo_url,omitempty"`

	// Language preference
	PreferredLanguage string `gorm:"size:50;default:'bn'" json:"preferred_language"` // bn (Bangla), en (English), roh (Rohingya)

	// Deceased status
	DeceasedBoolean  bool       `gorm:"default:false" json:"deceased_boolean"`
	DeceasedDateTime *time.Time `json:"deceased_datetime,omitempty"`

	// Relationships
	Encounters    []Encounter    `gorm:"foreignKey:PatientID" json:"encounters,omitempty"`
	Appointments  []Appointment  `gorm:"foreignKey:PatientID" json:"appointments,omitempty"`
	Prescriptions []Prescription `gorm:"foreignKey:PatientID" json:"prescriptions,omitempty"`
	LabOrders     []LabOrder     `gorm:"foreignKey:PatientID" json:"lab_orders,omitempty"`

	// Clinical Decision Support
	Allergies []string `gorm:"type:text[]" json:"allergies,omitempty"`
}

// TableName overrides the table name
func (Patient) TableName() string {
	return "patients"
}

// IsRohingya checks if the patient is a Rohingya refugee
func (p *Patient) IsRohingya() bool {
	return p.Nationality == "rohingya"
}

// IsBangladeshi checks if the patient is a Bangladeshi citizen
func (p *Patient) IsBangladeshi() bool {
	return p.Nationality == "bangladeshi"
}

// GetFullName returns the patient's full name
func (p *Patient) GetFullName() string {
	if p.MiddleName != "" {
		return p.GivenName + " " + p.MiddleName + " " + p.FamilyName
	}
	return p.GivenName + " " + p.FamilyName
}

// GetAge calculates the patient's age in years
func (p *Patient) GetAge() int {
	if p.BirthDate == nil {
		return 0
	}
	now := time.Now()
	age := now.Year() - p.BirthDate.Year()
	if now.YearDay() < p.BirthDate.YearDay() {
		age--
	}
	return age
}

// Validate performs validation based on nationality
func (p *Patient) Validate() error {
	if p.Nationality == "" {
		return ErrNationalityRequired
	}

	// Validate Bangladeshi-specific fields
	if p.IsBangladeshi() {
		if p.NationalID == "" && p.BirthRegNo == "" {
			return ErrBangladeshiIDRequired
		}
	}

	// Validate Rohingya-specific fields
	if p.IsRohingya() {
		if p.UNHCRNumber == "" {
			return ErrRohingyaUNHCRRequired
		}
	}

	return nil
}

// Custom errors
var (
	ErrNationalityRequired   = &ValidationError{Field: "nationality", Message: "Nationality is required"}
	ErrBangladeshiIDRequired = &ValidationError{Field: "national_id", Message: "National ID or Birth Registration Number is required for Bangladeshi citizens"}
	ErrRohingyaUNHCRRequired = &ValidationError{Field: "unhcr_number", Message: "UNHCR number is required for Rohingya refugees"}
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}
