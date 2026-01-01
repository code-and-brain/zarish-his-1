package service

import (
	"errors"
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type PharmacyService struct {
	repo *repository.PharmacyRepository
}

func NewPharmacyService(repo *repository.PharmacyRepository) *PharmacyService {
	return &PharmacyService{repo: repo}
}

func (s *PharmacyService) AddStock(stock *models.PharmacyStock) error {
	// Validate expiry date
	if stock.ExpiryDate.Before(time.Now()) {
		return errors.New("cannot add expired medication to stock")
	}

	return s.repo.AddStock(stock)
}

func (s *PharmacyService) GetAvailableStock(medicationID uint) ([]models.PharmacyStock, error) {
	return s.repo.GetStock(medicationID)
}

func (s *PharmacyService) GetLowStockAlerts() ([]models.PharmacyStock, error) {
	return s.repo.GetLowStock(10)
}

func (s *PharmacyService) DispenseMedication(dispensing *models.Dispensing) error {
	// Check if sufficient stock available
	stocks, err := s.repo.GetStock(dispensing.MedicationID)
	if err != nil {
		return err
	}

	totalAvailable := 0
	for _, stock := range stocks {
		totalAvailable += stock.Quantity
	}

	if totalAvailable < dispensing.QuantityDispensed {
		return errors.New("insufficient stock available")
	}

	dispensing.DispensedAt = time.Now()
	dispensing.Status = "dispensed"

	return s.repo.CreateDispensing(dispensing)
}

func (s *PharmacyService) GetDispensingQueue() ([]models.Prescription, error) {
	return s.repo.GetPendingPrescriptions()
}

func (s *PharmacyService) GetPatientDispensingHistory(patientID uint) ([]models.Dispensing, error) {
	return s.repo.GetDispensingHistory(patientID)
}

func (s *PharmacyService) GetStockMovementReport(medicationID uint, startDate, endDate time.Time) ([]models.StockMovement, error) {
	return s.repo.GetStockMovements(medicationID, startDate, endDate)
}
