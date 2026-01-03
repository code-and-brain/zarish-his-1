package integration

import (
	"time"

	"zarish-his/backend/internal/domain/models"
	"zarish-his/backend/internal/repository/postgres"
)

type BillingService struct {
	repo *postgres.BillingRepository
}

func NewBillingService(repo *postgres.BillingRepository) *BillingService {
	return &BillingService{repo: repo}
}

func (s *BillingService) GenerateInvoice(invoice *models.Invoice) error {
	return s.repo.CreateInvoice(invoice)
}

func (s *BillingService) GetInvoice(id uint) (*models.Invoice, error) {
	return s.repo.GetInvoice(id)
}

func (s *BillingService) GetPatientInvoices(patientID uint) ([]models.Invoice, error) {
	return s.repo.GetPatientInvoices(patientID)
}

func (s *BillingService) GetOutstandingInvoices() ([]models.Invoice, error) {
	return s.repo.GetOutstandingInvoices()
}

func (s *BillingService) RecordPayment(payment *models.Payment) error {
	return s.repo.CreatePayment(payment)
}

func (s *BillingService) GetInvoicePayments(invoiceID uint) ([]models.Payment, error) {
	return s.repo.GetInvoicePayments(invoiceID)
}

func (s *BillingService) GetPaymentReport(startDate, endDate time.Time) ([]models.Payment, error) {
	return s.repo.GetPaymentReport(startDate, endDate)
}

func (s *BillingService) SubmitInsuranceClaim(claim *models.InsuranceClaim) error {
	return s.repo.CreateClaim(claim)
}

func (s *BillingService) GetClaim(id uint) (*models.InsuranceClaim, error) {
	return s.repo.GetClaim(id)
}

func (s *BillingService) GetPendingClaims() ([]models.InsuranceClaim, error) {
	return s.repo.GetPendingClaims()
}

func (s *BillingService) ApproveClaim(id uint, approvedAmount float64) error {
	return s.repo.UpdateClaimStatus(id, "Approved", approvedAmount, "")
}

func (s *BillingService) RejectClaim(id uint, reason string) error {
	return s.repo.UpdateClaimStatus(id, "Rejected", 0, reason)
}
