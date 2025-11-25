package repository

import (
	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) CreatePatient(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *PatientRepository) GetPatientByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	return &patient, err
}

func (r *PatientRepository) UpdatePatient(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *PatientRepository) ListPatients(offset, limit int) ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.Offset(offset).Limit(limit).Find(&patients).Error
	return patients, err
}
