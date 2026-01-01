package repository

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type PharmacyRepository struct {
	db *gorm.DB
}

func NewPharmacyRepository(db *gorm.DB) *PharmacyRepository {
	return &PharmacyRepository{db: db}
}

// Stock Operations
func (r *PharmacyRepository) AddStock(stock *models.PharmacyStock) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(stock).Error; err != nil {
			return err
		}

		// Record stock movement
		movement := &models.StockMovement{
			Type:         "purchase",
			MedicationID: stock.MedicationID,
			Quantity:     stock.Quantity,
			BatchNumber:  stock.BatchNumber,
			Reference:    "STOCK-ADD",
			PerformedAt:  time.Now(),
		}
		return tx.Create(movement).Error
	})
}

func (r *PharmacyRepository) GetStock(medicationID uint) ([]models.PharmacyStock, error) {
	var stocks []models.PharmacyStock
	err := r.db.Where("medication_id = ? AND quantity > 0", medicationID).
		Preload("Medication").
		Order("expiry_date ASC").
		Find(&stocks).Error
	return stocks, err
}

func (r *PharmacyRepository) GetLowStock(threshold int) ([]models.PharmacyStock, error) {
	var stocks []models.PharmacyStock
	err := r.db.Preload("Medication").
		Where("quantity <= reorder_level").
		Find(&stocks).Error
	return stocks, err
}

func (r *PharmacyRepository) UpdateStockQuantity(id uint, quantity int) error {
	return r.db.Model(&models.PharmacyStock{}).
		Where("id = ?", id).
		Update("quantity", quantity).Error
}

// Dispensing Operations
func (r *PharmacyRepository) CreateDispensing(dispensing *models.Dispensing) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create dispensing record
		if err := tx.Create(dispensing).Error; err != nil {
			return err
		}

		// Deduct from stock (FIFO - First Expiry First Out)
		var stock models.PharmacyStock
		if err := tx.Where("medication_id = ? AND quantity >= ?",
			dispensing.MedicationID, dispensing.QuantityDispensed).
			Order("expiry_date ASC").
			First(&stock).Error; err != nil {
			return err
		}

		newQty := stock.Quantity - dispensing.QuantityDispensed
		if err := tx.Model(&stock).Update("quantity", newQty).Error; err != nil {
			return err
		}

		// Record stock movement
		movement := &models.StockMovement{
			Type:         "dispensing",
			MedicationID: dispensing.MedicationID,
			Quantity:     -dispensing.QuantityDispensed,
			BatchNumber:  stock.BatchNumber,
			Reference:    "DISP-" + string(rune(dispensing.ID)),
			PerformedBy:  dispensing.DispensedBy,
			PerformedAt:  dispensing.DispensedAt,
		}
		return tx.Create(movement).Error
	})
}

func (r *PharmacyRepository) GetPendingPrescriptions() ([]models.Prescription, error) {
	var prescriptions []models.Prescription
	err := r.db.Preload("Patient").
		Preload("Medication").
		Where("status = ?", "active").
		Where("id NOT IN (SELECT prescription_id FROM dispensing)").
		Find(&prescriptions).Error
	return prescriptions, err
}

func (r *PharmacyRepository) GetDispensingHistory(patientID uint) ([]models.Dispensing, error) {
	var dispensing []models.Dispensing
	err := r.db.Preload("Medication").
		Preload("Prescription").
		Where("patient_id = ?", patientID).
		Order("dispensed_at DESC").
		Find(&dispensing).Error
	return dispensing, err
}

// Stock Movement
func (r *PharmacyRepository) GetStockMovements(medicationID uint, startDate, endDate time.Time) ([]models.StockMovement, error) {
	var movements []models.StockMovement
	query := r.db.Preload("Medication").Where("medication_id = ?", medicationID)

	if !startDate.IsZero() {
		query = query.Where("performed_at >= ?", startDate)
	}
	if !endDate.IsZero() {
		query = query.Where("performed_at <= ?", endDate)
	}

	err := query.Order("performed_at DESC").Find(&movements).Error
	return movements, err
}
