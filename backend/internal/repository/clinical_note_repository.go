package repository

import (
	"errors"

	"github.com/ZarishSphere-Platform/zarish-his/internal/models"
	"gorm.io/gorm"
)

type ClinicalNoteRepository struct {
	db *gorm.DB
}

func NewClinicalNoteRepository(db *gorm.DB) *ClinicalNoteRepository {
	return &ClinicalNoteRepository{db: db}
}

func (r *ClinicalNoteRepository) Create(note *models.ClinicalNote) (*models.ClinicalNote, error) {
	if err := r.db.Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (r *ClinicalNoteRepository) FindByID(id uint) (*models.ClinicalNote, error) {
	var note models.ClinicalNote
	if err := r.db.First(&note, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &note, nil
}

func (r *ClinicalNoteRepository) Update(note *models.ClinicalNote) (*models.ClinicalNote, error) {
	if err := r.db.Save(note).Error; err != nil {
		return nil, err
	}
	return note, nil
}

func (r *ClinicalNoteRepository) ListByEncounter(encounterID uint) ([]*models.ClinicalNote, error) {
	var notes []*models.ClinicalNote
	if err := r.db.Where("encounter_id = ?", encounterID).Order("note_date DESC").Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *ClinicalNoteRepository) ListByPatient(patientID uint, limit int) ([]*models.ClinicalNote, error) {
	var notes []*models.ClinicalNote
	if err := r.db.Where("patient_id = ?", patientID).Order("note_date DESC").Limit(limit).Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}
