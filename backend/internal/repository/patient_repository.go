package repository

import (
	"errors"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("record not found")

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) Create(patient *models.Patient) (*models.Patient, error) {
	if err := r.db.Create(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *PatientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.First(&patient, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) FindByIDWithRelations(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.Preload("Encounters").Preload("Appointments").
		Preload("Prescriptions").Preload("LabOrders").
		First(&patient, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &patient, nil
}

func (r *PatientRepository) Update(patient *models.Patient) (*models.Patient, error) {
	if err := r.db.Save(patient).Error; err != nil {
		return nil, err
	}
	return patient, nil
}

func (r *PatientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}

func (r *PatientRepository) List(offset, limit int, nationality, search string) ([]*models.Patient, int64, error) {
	var patients []*models.Patient
	var total int64

	query := r.db.Model(&models.Patient{})

	// Filter by nationality
	if nationality != "" {
		query = query.Where("nationality = ?", nationality)
	}

	// Search filter
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where(
			"given_name ILIKE ? OR family_name ILIKE ? OR mrn ILIKE ? OR phone ILIKE ? OR email ILIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, searchPattern,
		)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&patients).Error; err != nil {
		return nil, 0, err
	}

	return patients, total, nil
}

func (r *PatientRepository) Search(query string) ([]*models.Patient, error) {
	var patients []*models.Patient
	searchPattern := "%" + query + "%"

	if err := r.db.Where(
		"given_name ILIKE ? OR family_name ILIKE ? OR mrn = ? OR national_id = ? OR unhcr_number = ? OR phone ILIKE ?",
		searchPattern, searchPattern, query, query, query, searchPattern,
	).Limit(20).Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}

func (r *PatientRepository) GetLastPatient() (*models.Patient, error) {
	var patient models.Patient
	if err := r.db.Order("id DESC").First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &patient, nil
}
