package repository

import (
	"errors"

	"github.com/ZarishSphere-Platform/zarish-his/internal/models"
	"gorm.io/gorm"
)

type EncounterRepository struct {
	db *gorm.DB
}

func NewEncounterRepository(db *gorm.DB) *EncounterRepository {
	return &EncounterRepository{db: db}
}

func (r *EncounterRepository) Create(encounter *models.Encounter) (*models.Encounter, error) {
	if err := r.db.Create(encounter).Error; err != nil {
		return nil, err
	}
	return encounter, nil
}

func (r *EncounterRepository) FindByID(id uint) (*models.Encounter, error) {
	var encounter models.Encounter
	if err := r.db.Preload("Patient").Preload("VitalSigns").Preload("ClinicalNotes").
		Preload("Prescriptions").Preload("LabOrders").
		First(&encounter, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &encounter, nil
}

func (r *EncounterRepository) Update(encounter *models.Encounter) (*models.Encounter, error) {
	if err := r.db.Save(encounter).Error; err != nil {
		return nil, err
	}
	return encounter, nil
}

func (r *EncounterRepository) ListByPatient(patientID uint, offset, limit int) ([]*models.Encounter, int64, error) {
	var encounters []*models.Encounter
	var total int64

	query := r.db.Model(&models.Encounter{}).Where("patient_id = ?", patientID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("period_start DESC").Find(&encounters).Error; err != nil {
		return nil, 0, err
	}

	return encounters, total, nil
}
