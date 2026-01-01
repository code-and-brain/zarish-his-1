package repository

import (
	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type RadiologyRepository struct {
	db *gorm.DB
}

func NewRadiologyRepository(db *gorm.DB) *RadiologyRepository {
	return &RadiologyRepository{db: db}
}

// Study Operations

func (r *RadiologyRepository) CreateStudy(study *models.ImagingStudy) error {
	return r.db.Create(study).Error
}

func (r *RadiologyRepository) GetStudy(id uint) (*models.ImagingStudy, error) {
	var study models.ImagingStudy
	err := r.db.Preload("Patient").Preload("Series.Instances").Preload("Report").First(&study, id).Error
	return &study, err
}

func (r *RadiologyRepository) GetStudyByAccession(accession string) (*models.ImagingStudy, error) {
	var study models.ImagingStudy
	err := r.db.Preload("Patient").Preload("Series").First(&study, "accession_number = ?", accession).Error
	return &study, err
}

func (r *RadiologyRepository) ListStudies(patientID uint, status string, limit, offset int) ([]models.ImagingStudy, int64, error) {
	var studies []models.ImagingStudy
	var total int64

	query := r.db.Model(&models.ImagingStudy{})

	if patientID > 0 {
		query = query.Where("patient_id = ?", patientID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Preload("Patient").Preload("Report").Order("created_at desc").Limit(limit).Offset(offset).Find(&studies).Error
	return studies, total, err
}

func (r *RadiologyRepository) UpdateStudyStatus(id uint, status string) error {
	return r.db.Model(&models.ImagingStudy{}).Where("id = ?", id).Update("status", status).Error
}

// Report Operations

func (r *RadiologyRepository) CreateReport(report *models.RadiologyReport) error {
	return r.db.Create(report).Error
}

func (r *RadiologyRepository) UpdateReport(report *models.RadiologyReport) error {
	return r.db.Save(report).Error
}

func (r *RadiologyRepository) GetReportByStudyID(studyID uint) (*models.RadiologyReport, error) {
	var report models.RadiologyReport
	err := r.db.Where("study_id = ?", studyID).First(&report).Error
	return &report, err
}

// Worklist Operations

func (r *RadiologyRepository) GetWorklist(status []string) ([]models.ImagingStudy, error) {
	var studies []models.ImagingStudy
	err := r.db.Preload("Patient").Where("status IN ?", status).Order("created_at asc").Find(&studies).Error
	return studies, err
}
