package models

import (
	"time"

	"gorm.io/gorm"
)

// PharmacyStock tracks medication inventory
type PharmacyStock struct {
	gorm.Model
	MedicationID uint       `json:"medication_id" gorm:"index;not null"`
	Medication   Medication `json:"medication,omitempty" gorm:"foreignKey:MedicationID"`
	Quantity     int        `json:"quantity" gorm:"not null"`
	BatchNumber  string     `json:"batch_number" gorm:"size:100"`
	ExpiryDate   time.Time  `json:"expiry_date"`
	Location     string     `json:"location" gorm:"size:100"` // e.g., "Shelf A-12"
	CostPrice    float64    `json:"cost_price"`
	SellingPrice float64    `json:"selling_price"`
	ReorderLevel int        `json:"reorder_level" gorm:"default:10"`
	Notes        string     `json:"notes" gorm:"type:text"`
}

// Dispensing records medication dispensed to patients
type Dispensing struct {
	gorm.Model
	PrescriptionID    uint         `json:"prescription_id" gorm:"index;not null"`
	Prescription      Prescription `json:"prescription,omitempty" gorm:"foreignKey:PrescriptionID"`
	PatientID         uint         `json:"patient_id" gorm:"index;not null"`
	Patient           Patient      `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	MedicationID      uint         `json:"medication_id" gorm:"index;not null"`
	Medication        Medication   `json:"medication,omitempty" gorm:"foreignKey:MedicationID"`
	QuantityDispensed int          `json:"quantity_dispensed" gorm:"not null"`
	BatchNumber       string       `json:"batch_number" gorm:"size:100"`
	DispensedBy       uint         `json:"dispensed_by"` // User/Pharmacist ID
	DispensedAt       time.Time    `json:"dispensed_at" gorm:"not null"`
	Instructions      string       `json:"instructions" gorm:"type:text"`
	Notes             string       `json:"notes" gorm:"type:text"`
	Status            string       `json:"status" gorm:"size:50;default:'dispensed'"` // dispensed, returned
}

// StockMovement tracks all stock in/out transactions
type StockMovement struct {
	gorm.Model
	Type         string     `json:"type" gorm:"size:50;not null"` // purchase, dispensing, adjustment, return, expired
	MedicationID uint       `json:"medication_id" gorm:"index;not null"`
	Medication   Medication `json:"medication,omitempty" gorm:"foreignKey:MedicationID"`
	Quantity     int        `json:"quantity" gorm:"not null"` // positive for in, negative for out
	BatchNumber  string     `json:"batch_number" gorm:"size:100"`
	Reference    string     `json:"reference" gorm:"size:200"` // e.g., "PO-12345", "Dispensing-67"
	Reason       string     `json:"reason" gorm:"type:text"`
	PerformedBy  uint       `json:"performed_by"`
	PerformedAt  time.Time  `json:"performed_at" gorm:"not null"`
}

// TableName overrides
func (PharmacyStock) TableName() string {
	return "pharmacy_stock"
}

func (Dispensing) TableName() string {
	return "dispensing"
}

func (StockMovement) TableName() string {
	return "stock_movements"
}
