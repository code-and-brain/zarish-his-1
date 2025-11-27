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
