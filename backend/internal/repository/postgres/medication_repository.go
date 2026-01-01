package repository

import (
	"errors"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type MedicationRepository struct {
	db *gorm.DB
}

func NewMedicationRepository(db *gorm.DB) *MedicationRepository {
	return &MedicationRepository{db: db}
}

// Medication methods
func (r *MedicationRepository) CreateMedication(med *models.Medication) (*models.Medication, error) {
	if err := r.db.Create(med).Error; err != nil {
		return nil, err
	}
	return med, nil
}

func (r *MedicationRepository) FindMedicationByID(id uint) (*models.Medication, error) {
	var med models.Medication
	if err := r.db.First(&med, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &med, nil
}

func (r *MedicationRepository) SearchMedications(query string) ([]*models.Medication, error) {
	var meds []*models.Medication
	searchPattern := "%" + query + "%"
	if err := r.db.Where("name ILIKE ? OR generic_name ILIKE ? OR brand_name ILIKE ?", searchPattern, searchPattern, searchPattern).Find(&meds).Error; err != nil {
		return nil, err
	}
	return meds, nil
}

// Prescription methods
func (r *MedicationRepository) CreatePrescription(prescription *models.Prescription) (*models.Prescription, error) {
	if err := r.db.Create(prescription).Error; err != nil {
		return nil, err
	}
	return prescription, nil
}

func (r *MedicationRepository) FindPrescriptionByID(id uint) (*models.Prescription, error) {
	var prescription models.Prescription
	if err := r.db.Preload("Medication").Preload("Patient").First(&prescription, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &prescription, nil
}

func (r *MedicationRepository) UpdatePrescription(prescription *models.Prescription) (*models.Prescription, error) {
	if err := r.db.Save(prescription).Error; err != nil {
		return nil, err
	}
	return prescription, nil
}

func (r *MedicationRepository) ListPrescriptionsByPatient(patientID uint) ([]*models.Prescription, error) {
	var prescriptions []*models.Prescription
	if err := r.db.Preload("Medication").Where("patient_id = ?", patientID).Order("start_date DESC").Find(&prescriptions).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}

func (r *MedicationRepository) ListActivePrescriptions(patientID uint) ([]*models.Prescription, error) {
	var prescriptions []*models.Prescription
	if err := r.db.Preload("Medication").Where("patient_id = ? AND status = 'active'", patientID).Order("start_date DESC").Find(&prescriptions).Error; err != nil {
		return nil, err
	}
	return prescriptions, nil
}
