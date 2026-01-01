package models

import (
	"time"

	"gorm.io/gorm"
)

// Invoice represents a patient bill
type Invoice struct {
	gorm.Model
	PatientID     uint          `json:"patient_id" gorm:"index;not null"`
	Patient       Patient       `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	EncounterID   *uint         `json:"encounter_id" gorm:"index"`
	Encounter     *Encounter    `json:"encounter,omitempty" gorm:"foreignKey:EncounterID"`
	InvoiceNumber string        `json:"invoice_number" gorm:"size:50;unique;not null"`
	InvoiceDate   time.Time     `json:"invoice_date" gorm:"not null"`
	DueDate       time.Time     `json:"due_date"`
	TotalAmount   float64       `json:"total_amount" gorm:"not null"`
	PaidAmount    float64       `json:"paid_amount" gorm:"default:0"`
	BalanceAmount float64       `json:"balance_amount" gorm:"default:0"`
	Status        string        `json:"status" gorm:"size:50;default:'pending'"` // pending, partial, paid, cancelled
	Notes         string        `json:"notes" gorm:"type:text"`
	Items         []InvoiceItem `json:"items,omitempty" gorm:"foreignKey:InvoiceID"`
	Payments      []Payment     `json:"payments,omitempty" gorm:"foreignKey:InvoiceID"`
}

// InvoiceItem represents a line item on an invoice
type InvoiceItem struct {
	gorm.Model
	InvoiceID   uint    `json:"invoice_id" gorm:"index;not null"`
	ServiceCode string  `json:"service_code" gorm:"size:50"`
	Description string  `json:"description" gorm:"not null"`
	Quantity    int     `json:"quantity" gorm:"default:1"`
	UnitPrice   float64 `json:"unit_price" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Discount    float64 `json:"discount" gorm:"default:0"`
	Tax         float64 `json:"tax" gorm:"default:0"`
	NetAmount   float64 `json:"net_amount" gorm:"not null"`
}

// Payment represents a payment transaction
type Payment struct {
	gorm.Model
	InvoiceID     uint      `json:"invoice_id" gorm:"index;not null"`
	Invoice       Invoice   `json:"invoice,omitempty" gorm:"foreignKey:InvoiceID"`
	Amount        float64   `json:"amount" gorm:"not null"`
	PaymentMethod string    `json:"payment_method" gorm:"size:50;not null"` // cash, card, insurance, mobile_money
	TransactionID string    `json:"transaction_id" gorm:"size:100"`
	PaidAt        time.Time `json:"paid_at" gorm:"not null"`
	ReceivedBy    uint      `json:"received_by"` // User ID
	Notes         string    `json:"notes" gorm:"type:text"`
	Status        string    `json:"status" gorm:"size:50;default:'completed'"` // completed, pending, failed, refunded
}

// InsuranceClaim represents an insurance claim
type InsuranceClaim struct {
	gorm.Model
	PatientID         uint       `json:"patient_id" gorm:"index;not null"`
	Patient           Patient    `json:"patient,omitempty" gorm:"foreignKey:PatientID"`
	EncounterID       *uint      `json:"encounter_id" gorm:"index"`
	Encounter         *Encounter `json:"encounter,omitempty" gorm:"foreignKey:EncounterID"`
	InvoiceID         *uint      `json:"invoice_id" gorm:"index"`
	Invoice           *Invoice   `json:"invoice,omitempty" gorm:"foreignKey:InvoiceID"`
	ClaimNumber       string     `json:"claim_number" gorm:"size:50;unique;not null"`
	InsuranceProvider string     `json:"insurance_provider" gorm:"size:200;not null"`
	PolicyNumber      string     `json:"policy_number" gorm:"size:100"`
	ClaimAmount       float64    `json:"claim_amount" gorm:"not null"`
	ApprovedAmount    float64    `json:"approved_amount" gorm:"default:0"`
	Status            string     `json:"status" gorm:"size:50;default:'submitted'"` // submitted, under_review, approved, rejected, paid
	SubmittedAt       time.Time  `json:"submitted_at" gorm:"not null"`
	ReviewedAt        *time.Time `json:"reviewed_at"`
	PaidAt            *time.Time `json:"paid_at"`
	RejectionReason   string     `json:"rejection_reason" gorm:"type:text"`
	Notes             string     `json:"notes" gorm:"type:text"`
}

// TableName overrides
func (Invoice) TableName() string {
	return "invoices"
}

func (InvoiceItem) TableName() string {
	return "invoice_items"
}

func (Payment) TableName() string {
	return "payments"
}

func (InsuranceClaim) TableName() string {
	return "insurance_claims"
}
