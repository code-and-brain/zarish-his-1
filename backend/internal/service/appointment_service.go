package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type AppointmentService struct {
	repo *repository.AppointmentRepository
}

func NewAppointmentService(repo *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

func (s *AppointmentService) CreateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	if appointment.Status == "" {
		appointment.Status = "scheduled"
	}
	return s.repo.Create(appointment)
}

func (s *AppointmentService) GetAppointmentByID(id uint) (*models.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *AppointmentService) UpdateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	return s.repo.Update(appointment)
}

func (s *AppointmentService) CancelAppointment(id uint, reason string, cancelledBy uint) (*models.Appointment, error) {
	appointment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	appointment.Cancel(reason, cancelledBy)
	return s.repo.Update(appointment)
}

func (s *AppointmentService) ListAppointmentsByDate(date time.Time) ([]*models.Appointment, error) {
	start := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	end := start.Add(24 * time.Hour)
	return s.repo.ListByDateRange(start, end)
}

func (s *AppointmentService) ListPatientAppointments(patientID uint) ([]*models.Appointment, error) {
	return s.repo.ListByPatient(patientID)
}
