package repository

import (
	"errors"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type VitalSignsRepository struct {
	db *gorm.DB
}

func NewVitalSignsRepository(db *gorm.DB) *VitalSignsRepository {
	return &VitalSignsRepository{db: db}
}

func (r *VitalSignsRepository) Create(vitals *models.VitalSigns) (*models.VitalSigns, error) {
	if err := r.db.Create(vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

func (r *VitalSignsRepository) FindByID(id uint) (*models.VitalSigns, error) {
	var vitals models.VitalSigns
	if err := r.db.First(&vitals, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &vitals, nil
}

func (r *VitalSignsRepository) ListByEncounter(encounterID uint) ([]*models.VitalSigns, error) {
	var vitals []*models.VitalSigns
	if err := r.db.Where("encounter_id = ?", encounterID).Order("measured_at DESC").Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}

func (r *VitalSignsRepository) ListByPatient(patientID uint, limit int) ([]*models.VitalSigns, error) {
	var vitals []*models.VitalSigns
	if err := r.db.Where("patient_id = ?", patientID).Order("measured_at DESC").Limit(limit).Find(&vitals).Error; err != nil {
		return nil, err
	}
	return vitals, nil
}
