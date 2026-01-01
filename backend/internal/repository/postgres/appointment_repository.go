package repository

import (
	"errors"
	"time"

	"github.com/code-and-brain/zarish-his-1/backend/internal/models"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) (*models.Appointment, error) {
	if err := r.db.Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *AppointmentRepository) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	if err := r.db.Preload("Patient").First(&appointment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) (*models.Appointment, error) {
	if err := r.db.Save(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *AppointmentRepository) ListByDateRange(start, end time.Time) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.db.Preload("Patient").Where("scheduled_start BETWEEN ? AND ?", start, end).Order("scheduled_start ASC").Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) ListByPatient(patientID uint) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.db.Where("patient_id = ?", patientID).Order("scheduled_start DESC").Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}
