package service

import (
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type ReportingService struct {
	db *gorm.DB
}

func NewReportingService(db *gorm.DB) *ReportingService {
	return &ReportingService{db: db}
}

type DailyOPDStats struct {
	Date            string           `json:"date"`
	TotalEncounters int64            `json:"total_encounters"`
	ByServiceType   map[string]int64 `json:"by_service_type"`
	ByCategory      map[string]int64 `json:"by_category"`
}

func (s *ReportingService) GenerateDailyOPDReport(date time.Time) (*DailyOPDStats, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	var total int64
	if err := s.db.Model(&models.Encounter{}).
		Where("period_start >= ? AND period_start < ?", startOfDay, endOfDay).
		Count(&total).Error; err != nil {
		return nil, err
	}

	// Group by Service Type
	type Result struct {
		Key   string
		Count int64
	}

	var typeResults []Result
	if err := s.db.Model(&models.Encounter{}).
		Select("service_type as key, count(*) as count").
		Where("period_start >= ? AND period_start < ?", startOfDay, endOfDay).
		Group("service_type").
		Scan(&typeResults).Error; err != nil {
		return nil, err
	}

	byType := make(map[string]int64)
	for _, r := range typeResults {
		if r.Key == "" {
			r.Key = "Unspecified"
		}
		byType[r.Key] = r.Count
	}

	// Group by Service Category
	var catResults []Result
	if err := s.db.Model(&models.Encounter{}).
		Select("service_category as key, count(*) as count").
		Where("period_start >= ? AND period_start < ?", startOfDay, endOfDay).
		Group("service_category").
		Scan(&catResults).Error; err != nil {
		return nil, err
	}

	byCategory := make(map[string]int64)
	for _, r := range catResults {
		if r.Key == "" {
			r.Key = "Unspecified"
		}
		byCategory[r.Key] = r.Count
	}

	return &DailyOPDStats{
		Date:            date.Format("2006-01-02"),
		TotalEncounters: total,
		ByServiceType:   byType,
		ByCategory:      byCategory,
	}, nil
}

type DiseaseStats struct {
	Diagnosis string `json:"diagnosis"`
	Count     int64  `json:"count"`
}

func (s *ReportingService) GenerateDiseaseSurveillanceReport(startDate, endDate time.Time) ([]DiseaseStats, error) {
	var results []DiseaseStats

	// This is a simplified query. In a real system, we might need to parse JSON or split strings if multiple diagnoses are stored.
	// Assuming primary diagnosis is stored in 'diagnosis' field as a string.
	if err := s.db.Model(&models.Encounter{}).
		Select("diagnosis, count(*) as count").
		Where("period_start >= ? AND period_start < ? AND diagnosis != ''", startDate, endDate).
		Group("diagnosis").
		Order("count desc").
		Limit(20). // Top 20 diseases
		Scan(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
