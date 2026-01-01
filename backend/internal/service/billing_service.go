package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type BillingService struct {
	repo *repository.BillingRepository
}

func NewBillingService(repo *repository.BillingRepository) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) GenerateInvoice(invoice *models.Invoice) error {
	// Generate invoice number
	invoice.InvoiceNumber = s.generateInvoiceNumber()
	invoice.InvoiceDate = time.Now()

	// Set due date (30 days from invoice date)
	invoice.DueDate = invoice.InvoiceDate.AddDate(0, 0, 30)

	// Calculate item amounts
	for i := range invoice.Items {
		item := &invoice.Items[i]
		item.Amount = float64(item.Quantity) * item.UnitPrice
		item.NetAmount = item.Amount - item.Discount + item.Tax
	}

	invoice.Status = "pending"
	invoice.PaidAmount = 0

	return s.repo.CreateInvoice(invoice)
}

func (s *BillingService) GetInvoice(id uint) (*models.Invoice, error) {
	return s.repo.GetInvoice(id)
}

func (s *BillingService) GetPatientInvoices(patientID uint) ([]models.Invoice, error) {
	return s.repo.ListInvoices(patientID, "")
}

func (s *BillingService) GetOutstandingInvoices() ([]models.Invoice, error) {
	return s.repo.GetOutstandingInvoices()
}

func (s *BillingService) RecordPayment(payment *models.Payment) error {
	// Validate payment amount
	invoice, err := s.repo.GetInvoice(payment.InvoiceID)
	if err != nil {
		return err
	}

	if payment.Amount > invoice.BalanceAmount {
		return errors.New("payment amount exceeds balance")
	}

	payment.PaidAt = time.Now()
	payment.Status = "completed"

	return s.repo.CreatePayment(payment)
}

func (s *BillingService) GetInvoicePayments(invoiceID uint) ([]models.Payment, error) {
	return s.repo.GetPaymentsByInvoice(invoiceID)
}

func (s *BillingService) GetPaymentReport(startDate, endDate time.Time) ([]models.Payment, error) {
	return s.repo.GetPaymentsByDateRange(startDate, endDate)
}

func (s *BillingService) SubmitInsuranceClaim(claim *models.InsuranceClaim) error {
	// Generate claim number
	claim.ClaimNumber = s.generateClaimNumber()
	claim.SubmittedAt = time.Now()
	claim.Status = "submitted"

	return s.repo.CreateClaim(claim)
}

func (s *BillingService) GetClaim(id uint) (*models.InsuranceClaim, error) {
	return s.repo.GetClaim(id)
}

func (s *BillingService) GetPendingClaims() ([]models.InsuranceClaim, error) {
	return s.repo.ListClaims("submitted")
}

func (s *BillingService) ApproveClaim(id uint, approvedAmount float64) error {
	return s.repo.UpdateClaimStatus(id, "approved", approvedAmount, "")
}

func (s *BillingService) RejectClaim(id uint, reason string) error {
	return s.repo.UpdateClaimStatus(id, "rejected", 0, reason)
}

func (s *BillingService) MarkClaimPaid(id uint) error {
	claim, err := s.repo.GetClaim(id)
	if err != nil {
		return err
	}

	if claim.Status != "approved" {
		return errors.New("claim must be approved before marking as paid")
	}

	return s.repo.UpdateClaimStatus(id, "paid", claim.ApprovedAmount, "")
}

// Helper functions
func (s *BillingService) generateInvoiceNumber() string {
	return fmt.Sprintf("INV-%d-%d", time.Now().Unix(), time.Now().Nanosecond()%1000)
}

func (s *BillingService) generateClaimNumber() string {
	return fmt.Sprintf("CLM-%d-%d", time.Now().Unix(), time.Now().Nanosecond()%1000)
}
