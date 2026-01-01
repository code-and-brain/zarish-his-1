package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type RadiologyService struct {
	repo *repository.RadiologyRepository
}

func NewRadiologyService(repo *repository.RadiologyRepository) *RadiologyService {
	return &RadiologyService{repo: repo}
}

func (s *RadiologyService) CreateStudy(study *models.ImagingStudy) error {
	// Generate UIDs if not present (simplified for MVP)
	if study.StudyUID == "" {
		study.StudyUID = fmt.Sprintf("1.2.840.10008.%d", time.Now().UnixNano())
	}
	if study.AccessionNumber == "" {
		study.AccessionNumber = fmt.Sprintf("ACC-%d", time.Now().Unix())
	}

	study.Status = "scheduled"
	study.StartedAt = time.Now()

	return s.repo.CreateStudy(study)
}

func (s *RadiologyService) GetStudy(id uint) (*models.ImagingStudy, error) {
	return s.repo.GetStudy(id)
}

func (s *RadiologyService) ListStudies(patientID uint, status string, page, limit int) ([]models.ImagingStudy, int64, error) {
	offset := (page - 1) * limit
	return s.repo.ListStudies(patientID, status, limit, offset)
}

func (s *RadiologyService) StartExam(id uint, techID uint) error {
	study, err := s.repo.GetStudy(id)
	if err != nil {
		return err
	}

	if study.Status != "scheduled" {
		return errors.New("study can only be started from scheduled status")
	}

	study.Status = "in-progress"
	study.PerformingTechID = &techID

	// Use repository update or save
	// For simplicity, we'll use the status update method but ideally we'd update multiple fields
	return s.repo.UpdateStudyStatus(id, "in-progress")
}

func (s *RadiologyService) CompleteExam(id uint) error {
	// In a real system, we'd verify that images have been received
	// Also update CompletedAt in a real implementation
	return s.repo.UpdateStudyStatus(id, "completed")
}

func (s *RadiologyService) CreateReport(report *models.RadiologyReport) error {
	// Check if study exists
	study, err := s.repo.GetStudy(report.StudyID)
	if err != nil {
		return err
	}

	if study.Status != "completed" {
		return errors.New("cannot report on an incomplete study")
	}

	report.ReportedAt = time.Now()
	if report.Status == "" {
		report.Status = "draft"
	}

	return s.repo.CreateReport(report)
}

func (s *RadiologyService) FinalizeReport(id uint) error {
	// This would fetch the report, check permissions, and update status
	// For MVP, we'll assume the handler passes a report object with status="final" to UpdateReport
	return nil
}

func (s *RadiologyService) UpdateReport(report *models.RadiologyReport) error {
	if report.Status == "final" {
		now := time.Now()
		report.FinalizedAt = &now
	}
	return s.repo.UpdateReport(report)
}

func (s *RadiologyService) GetWorklist() ([]models.ImagingStudy, error) {
	// Radiologist worklist: completed studies needing reports
	// Technician worklist: scheduled studies
	// For now, return all active
	return s.repo.GetWorklist([]string{"scheduled", "in-progress", "completed"})
}
