package repository

import (
	"github.com/zarishsphere/zarish-his/internal/models"
	"gorm.io/gorm"
)

type ADTRepository struct {
	db *gorm.DB
}

func NewADTRepository(db *gorm.DB) *ADTRepository {
	return &ADTRepository{db: db}
}

// Ward Operations
func (r *ADTRepository) CreateWard(ward *models.Ward) error {
	return r.db.Create(ward).Error
}

func (r *ADTRepository) ListWards() ([]models.Ward, error) {
	var wards []models.Ward
	err := r.db.Preload("Rooms.Beds").Find(&wards).Error
	return wards, err
}

func (r *ADTRepository) GetWard(id uint) (*models.Ward, error) {
	var ward models.Ward
	err := r.db.Preload("Rooms.Beds").First(&ward, id).Error
	return &ward, err
}

// Room Operations
func (r *ADTRepository) CreateRoom(room *models.Room) error {
	return r.db.Create(room).Error
}

// Bed Operations
func (r *ADTRepository) CreateBed(bed *models.Bed) error {
	return r.db.Create(bed).Error
}

func (r *ADTRepository) UpdateBedStatus(id uint, status string) error {
	return r.db.Model(&models.Bed{}).Where("id = ?", id).Update("status", status).Error
}

func (r *ADTRepository) ListBeds(status string) ([]models.Bed, error) {
	var beds []models.Bed
	query := r.db.Model(&models.Bed{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&beds).Error
	return beds, err
}

// Admission Operations
func (r *ADTRepository) CreateAdmission(admission *models.Admission) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create admission record
		if err := tx.Create(admission).Error; err != nil {
			return err
		}
		// Update bed status to Occupied
		if err := tx.Model(&models.Bed{}).Where("id = ?", admission.BedID).Update("status", "Occupied").Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *ADTRepository) DischargePatient(admissionID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var admission models.Admission
		if err := tx.First(&admission, admissionID).Error; err != nil {
			return err
		}

		// Update admission status
		now := tx.NowFunc()
		if err := tx.Model(&admission).Updates(map[string]interface{}{
			"status":         "Discharged",
			"discharge_date": now,
		}).Error; err != nil {
			return err
		}

		// Update bed status to Available (or Cleaning)
		if err := tx.Model(&models.Bed{}).Where("id = ?", admission.BedID).Update("status", "Available").Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *ADTRepository) ListAdmissions(status string) ([]models.Admission, error) {
	var admissions []models.Admission
	query := r.db.Preload("Patient").Preload("Bed")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Find(&admissions).Error
	return admissions, err
}

func (r *ADTRepository) GetAdmission(id uint) (*models.Admission, error) {
	var admission models.Admission
	err := r.db.Preload("Patient").Preload("Bed").Preload("Ward").First(&admission, id).Error
	return &admission, err
}

// Transfer Operations
func (r *ADTRepository) CreateTransfer(transfer *models.Transfer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create transfer record
		if err := tx.Create(transfer).Error; err != nil {
			return err
		}

		// Update admission's current location
		if err := tx.Model(&models.Admission{}).Where("id = ?", transfer.AdmissionID).Updates(map[string]interface{}{
			"ward_id": transfer.ToWardID,
			"bed_id":  transfer.ToBedID,
		}).Error; err != nil {
			return err
		}

		// Update old bed status to Available
		if err := tx.Model(&models.Bed{}).Where("id = ?", transfer.FromBedID).Update("status", "Available").Error; err != nil {
			return err
		}

		// Update new bed status to Occupied
		if err := tx.Model(&models.Bed{}).Where("id = ?", transfer.ToBedID).Update("status", "Occupied").Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *ADTRepository) ListTransfers(admissionID uint) ([]models.Transfer, error) {
	var transfers []models.Transfer
	query := r.db.Preload("FromWard").Preload("FromBed").Preload("ToWard").Preload("ToBed")
	if admissionID > 0 {
		query = query.Where("admission_id = ?", admissionID)
	}
	err := query.Find(&transfers).Error
	return transfers, err
}

// Discharge Summary Operations
func (r *ADTRepository) CreateDischargeSummary(summary *models.DischargeSummary) error {
	return r.db.Create(summary).Error
}

func (r *ADTRepository) GetDischargeSummary(admissionID uint) (*models.DischargeSummary, error) {
	var summary models.DischargeSummary
	err := r.db.Where("admission_id = ?", admissionID).First(&summary).Error
	return &summary, err
}
