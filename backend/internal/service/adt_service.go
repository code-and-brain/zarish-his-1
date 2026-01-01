package service

import (
	"errors"
	"time"

	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
)

type ADTService struct {
	repo *repository.ADTRepository
}

func NewADTService(repo *repository.ADTRepository) *ADTService {
	return &ADTService{repo: repo}
}

func (s *ADTService) CreateWard(ward *models.Ward) error {
	return s.repo.CreateWard(ward)
}

func (s *ADTService) ListWards() ([]models.Ward, error) {
	return s.repo.ListWards()
}

func (s *ADTService) CreateRoom(room *models.Room) error {
	return s.repo.CreateRoom(room)
}

func (s *ADTService) CreateBed(bed *models.Bed) error {
	return s.repo.CreateBed(bed)
}

func (s *ADTService) ListBeds(status string) ([]models.Bed, error) {
	return s.repo.ListBeds(status)
}

func (s *ADTService) AdmitPatient(admission *models.Admission) error {
	// Validate bed availability
	beds, err := s.repo.ListBeds("Available")
	if err != nil {
		return err
	}

	isAvailable := false
	for _, bed := range beds {
		if bed.ID == admission.BedID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return errors.New("selected bed is not available")
	}

	admission.AdmissionDate = time.Now()
	admission.Status = "Admitted"
	return s.repo.CreateAdmission(admission)
}

func (s *ADTService) DischargePatient(admissionID uint) error {
	return s.repo.DischargePatient(admissionID)
}

func (s *ADTService) ListActiveAdmissions() ([]models.Admission, error) {
	return s.repo.ListAdmissions("Admitted")
}

func (s *ADTService) GetAdmission(id uint) (*models.Admission, error) {
	return s.repo.GetAdmission(id)
}

// TransferPatient transfers a patient to a new ward/bed
func (s *ADTService) TransferPatient(transfer *models.Transfer) error {
	// Validate destination bed is available
	beds, err := s.repo.ListBeds("Available")
	if err != nil {
		return err
	}

	isAvailable := false
	for _, bed := range beds {
		if bed.ID == transfer.ToBedID {
			isAvailable = true
			break
		}
	}

	if !isAvailable {
		return errors.New("destination bed is not available")
	}

	transfer.TransferDate = time.Now()
	return s.repo.CreateTransfer(transfer)
}

func (s *ADTService) ListTransfers(admissionID uint) ([]models.Transfer, error) {
	return s.repo.ListTransfers(admissionID)
}

// CreateDischargeSummary creates a discharge summary and updates admission status
func (s *ADTService) CreateDischargeSummary(summary *models.DischargeSummary) error {
	summary.DischargeDate = time.Now()

	// Create the summary
	if err := s.repo.CreateDischargeSummary(summary); err != nil {
		return err
	}

	// Discharge the patient
	return s.repo.DischargePatient(summary.AdmissionID)
}

func (s *ADTService) GetDischargeSummary(admissionID uint) (*models.DischargeSummary, error) {
	return s.repo.GetDischargeSummary(admissionID)
}
