package service

import (
	"time"

	"github.com/code-and-brain/zarish-his-1/backend/internal/models"
	"github.com/code-and-brain/zarish-his-1/backend/internal/repository"
)

type LabService struct {
	repo *repository.LabRepository
}

func NewLabService(repo *repository.LabRepository) *LabService {
	return &LabService{repo: repo}
}

func (s *LabService) CreateLabTest(test *models.LabTest) (*models.LabTest, error) {
	return s.repo.CreateLabTest(test)
}

func (s *LabService) ListLabTests() ([]*models.LabTest, error) {
	return s.repo.ListLabTests()
}

func (s *LabService) CreateLabOrder(order *models.LabOrder) (*models.LabOrder, error) {
	if order.OrderDate.IsZero() {
		order.OrderDate = time.Now()
	}
	if order.Status == "" {
		order.Status = "ordered"
	}
	return s.repo.CreateLabOrder(order)
}

func (s *LabService) GetLabOrderByID(id uint) (*models.LabOrder, error) {
	return s.repo.FindLabOrderByID(id)
}

func (s *LabService) ListPatientLabOrders(patientID uint) ([]*models.LabOrder, error) {
	return s.repo.ListLabOrdersByPatient(patientID)
}

func (s *LabService) AddLabResult(result *models.LabResult) (*models.LabResult, error) {
	// Fetch the test definition to check reference ranges
	test, err := s.repo.FindLabTestByID(result.LabTestID)
	if err == nil {
		// Auto-determine abnormal flag
		result.DetermineAbnormalFlag(test)
	}

	if result.ResultDate.IsZero() {
		result.ResultDate = time.Now()
	}

	return s.repo.CreateLabResult(result)
}
