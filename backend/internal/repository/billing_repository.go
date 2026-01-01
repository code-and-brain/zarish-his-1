package repository

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type BillingRepository struct {
	db *gorm.DB
}

func NewBillingRepository(db *gorm.DB) *BillingRepository {
	return &BillingRepository{db: db}
}

// Invoice Operations
func (r *BillingRepository) CreateInvoice(invoice *models.Invoice) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create invoice
		if err := tx.Create(invoice).Error; err != nil {
			return err
		}

		// Calculate totals from items
		var totalAmount float64
		for _, item := range invoice.Items {
			totalAmount += item.NetAmount
		}

		invoice.TotalAmount = totalAmount
		invoice.BalanceAmount = totalAmount - invoice.PaidAmount

		return tx.Save(invoice).Error
	})
}

func (r *BillingRepository) GetInvoice(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.Preload("Patient").
		Preload("Encounter").
		Preload("Items").
		Preload("Payments").
		First(&invoice, id).Error
	return &invoice, err
}

func (r *BillingRepository) ListInvoices(patientID uint, status string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	query := r.db.Preload("Patient").Preload("Items")

	if patientID > 0 {
		query = query.Where("patient_id = ?", patientID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Order("invoice_date DESC").Find(&invoices).Error
	return invoices, err
}

func (r *BillingRepository) GetOutstandingInvoices() ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := r.db.Preload("Patient").
		Where("status IN ?", []string{"pending", "partial"}).
		Where("balance_amount > 0").
		Order("due_date ASC").
		Find(&invoices).Error
	return invoices, err
}

func (r *BillingRepository) UpdateInvoiceStatus(id uint, status string) error {
	return r.db.Model(&models.Invoice{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// Payment Operations
func (r *BillingRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create payment
		if err := tx.Create(payment).Error; err != nil {
			return err
		}

		// Update invoice paid amount and balance
		var invoice models.Invoice
		if err := tx.First(&invoice, payment.InvoiceID).Error; err != nil {
			return err
		}

		invoice.PaidAmount += payment.Amount
		invoice.BalanceAmount = invoice.TotalAmount - invoice.PaidAmount

		// Update invoice status
		if invoice.BalanceAmount <= 0 {
			invoice.Status = "paid"
		} else if invoice.PaidAmount > 0 {
			invoice.Status = "partial"
		}

		return tx.Save(&invoice).Error
	})
}

func (r *BillingRepository) GetPaymentsByInvoice(invoiceID uint) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Where("invoice_id = ?", invoiceID).
		Order("paid_at DESC").
		Find(&payments).Error
	return payments, err
}

func (r *BillingRepository) GetPaymentsByDateRange(startDate, endDate time.Time) ([]models.Payment, error) {
	var payments []models.Payment
	err := r.db.Preload("Invoice").
		Where("paid_at >= ? AND paid_at <= ?", startDate, endDate).
		Order("paid_at DESC").
		Find(&payments).Error
	return payments, err
}

// Insurance Claim Operations
func (r *BillingRepository) CreateClaim(claim *models.InsuranceClaim) error {
	return r.db.Create(claim).Error
}

func (r *BillingRepository) GetClaim(id uint) (*models.InsuranceClaim, error) {
	var claim models.InsuranceClaim
	err := r.db.Preload("Patient").
		Preload("Encounter").
		Preload("Invoice").
		First(&claim, id).Error
	return &claim, err
}

func (r *BillingRepository) ListClaims(status string) ([]models.InsuranceClaim, error) {
	var claims []models.InsuranceClaim
	query := r.db.Preload("Patient").Preload("Invoice")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Order("submitted_at DESC").Find(&claims).Error
	return claims, err
}

func (r *BillingRepository) UpdateClaimStatus(id uint, status string, approvedAmount float64, reason string) error {
	updates := map[string]interface{}{
		"status":          status,
		"approved_amount": approvedAmount,
	}

	if status == "approved" || status == "rejected" {
		updates["reviewed_at"] = time.Now()
	}

	if status == "rejected" && reason != "" {
		updates["rejection_reason"] = reason
	}

	if status == "paid" {
		updates["paid_at"] = time.Now()
	}

	return r.db.Model(&models.InsuranceClaim{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *BillingRepository) GetClaimsByInsurer(insuranceProvider string) ([]models.InsuranceClaim, error) {
	var claims []models.InsuranceClaim
	err := r.db.Preload("Patient").
		Where("insurance_provider = ?", insuranceProvider).
		Order("submitted_at DESC").
		Find(&claims).Error
	return claims, err
}
