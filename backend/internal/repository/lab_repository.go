package repository

import (
	"errors"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type LabRepository struct {
	db *gorm.DB
}

func NewLabRepository(db *gorm.DB) *LabRepository {
	return &LabRepository{db: db}
}

// Lab Test methods
func (r *LabRepository) CreateLabTest(test *models.LabTest) (*models.LabTest, error) {
	if err := r.db.Create(test).Error; err != nil {
		return nil, err
	}
	return test, nil
}

func (r *LabRepository) ListLabTests() ([]*models.LabTest, error) {
	var tests []*models.LabTest
	if err := r.db.Where("active = ?", true).Order("name ASC").Find(&tests).Error; err != nil {
		return nil, err
	}
	return tests, nil
}

func (r *LabRepository) FindLabTestByID(id uint) (*models.LabTest, error) {
	var test models.LabTest
	if err := r.db.First(&test, id).Error; err != nil {
		return nil, err
	}
	return &test, nil
}

// Lab Order methods
func (r *LabRepository) CreateLabOrder(order *models.LabOrder) (*models.LabOrder, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *LabRepository) FindLabOrderByID(id uint) (*models.LabOrder, error) {
	var order models.LabOrder
	if err := r.db.Preload("Results.LabTest").Preload("Patient").First(&order, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &order, nil
}

func (r *LabRepository) ListLabOrdersByPatient(patientID uint) ([]*models.LabOrder, error) {
	var orders []*models.LabOrder
	if err := r.db.Preload("Results").Where("patient_id = ?", patientID).Order("order_date DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// Lab Result methods
func (r *LabRepository) CreateLabResult(result *models.LabResult) (*models.LabResult, error) {
	if err := r.db.Create(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *LabRepository) UpdateLabResult(result *models.LabResult) (*models.LabResult, error) {
	if err := r.db.Save(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
