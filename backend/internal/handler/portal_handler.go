package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type PortalHandler struct {
	patientService      *service.PatientService
	appointmentService  *service.AppointmentService
	labService          *service.LabService
	medicationService   *service.MedicationService
	clinicalNoteService *service.ClinicalNoteService
}

func NewPortalHandler(
	patientService *service.PatientService,
	appointmentService *service.AppointmentService,
	labService *service.LabService,
	medicationService *service.MedicationService,
	clinicalNoteService *service.ClinicalNoteService,
) *PortalHandler {
	return &PortalHandler{
		patientService:      patientService,
		appointmentService:  appointmentService,
		labService:          labService,
		medicationService:   medicationService,
		clinicalNoteService: clinicalNoteService,
	}
}

// GetDashboard returns a summary for the patient dashboard
func (h *PortalHandler) GetDashboard(c *gin.Context) {
	// In a real app, we'd get the patient ID from the authenticated user context
	// For MVP, we'll accept it as a query param or header, but strictly it should come from the token
	// c.GetInt("patient_id")

	patientID, _ := strconv.Atoi(c.Query("patient_id"))
	if patientID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient ID required"})
		return
	}

	patient, err := h.patientService.GetPatientByID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	appointments, err := h.appointmentService.ListPatientAppointments(uint(patientID))
	// Slice to get top 5
	if len(appointments) > 5 {
		appointments = appointments[:5]
	}

	c.JSON(http.StatusOK, gin.H{
		"patient":      patient,
		"appointments": appointments,
		"alerts":       []string{}, // Placeholder for alerts
	})
}

func (h *PortalHandler) GetAppointments(c *gin.Context) {
	patientID, _ := strconv.Atoi(c.Query("patient_id"))
	appointments, err := h.appointmentService.ListPatientAppointments(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointments, "total": len(appointments)})
}

func (h *PortalHandler) GetRecords(c *gin.Context) {
	patientID, _ := strconv.Atoi(c.Query("patient_id"))

	// Fetch various records
	// This is a simplified aggregation
	notes, err := h.clinicalNoteService.ListPatientNotes(uint(patientID), 10)
	if err != nil {
		notes = []*models.ClinicalNote{}
	}

	prescriptions, err := h.medicationService.ListPatientPrescriptions(uint(patientID))
	if err != nil {
		prescriptions = []*models.Prescription{}
	}

	labOrders, err := h.labService.ListPatientLabOrders(uint(patientID))
	if err != nil {
		labOrders = []*models.LabOrder{}
	}

	c.JSON(http.StatusOK, gin.H{
		"clinical_notes": notes,
		"prescriptions":  prescriptions,
		"lab_orders":     labOrders,
	})
}
