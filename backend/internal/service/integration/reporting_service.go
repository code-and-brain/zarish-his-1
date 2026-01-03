package integration

import (
	"time"

	"zarish-his/backend/internal/repository/postgres"
)

type ReportingService struct {
	repo *postgres.ReportingRepository
}

func NewReportingService(repo *postgres.ReportingRepository) *ReportingService {
	return &ReportingService{repo: repo}
}

func (s *ReportingService) GetPatientActivity(patientID uint, startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.repo.GetPatientActivity(patientID, startDate, endDate)
}

func (s *ReportingService) GetAppointmentAnalytics(startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.repo.GetAppointmentAnalytics(startDate, endDate)
}

func (s *ReportingService) GetBillingAnalytics(startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.repo.GetBillingAnalytics(startDate, endDate)
}

func (s *ReportingService) GetLabAnalytics(startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.repo.GetLabAnalytics(startDate, endDate)
}

func (s *ReportingService) GetPharmacyAnalytics(startDate, endDate time.Time) (map[string]interface{}, error) {
	return s.repo.GetPharmacyAnalytics(startDate, endDate)
}
