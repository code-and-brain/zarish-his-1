package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type MedicationHandler struct {
	service *service.MedicationService
}

func NewMedicationHandler(service *service.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: service}
}

func (h *MedicationHandler) CreateMedication(c *gin.Context) {
	var med models.Medication
	if err := c.ShouldBindJSON(&med); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMed, err := h.service.CreateMedication(&med)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMed)
}

func (h *MedicationHandler) SearchMedications(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	meds, err := h.service.SearchMedications(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, meds)
}

func (h *MedicationHandler) CreatePrescription(c *gin.Context) {
	var prescription models.Prescription
	if err := c.ShouldBindJSON(&prescription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPrescription, err := h.service.CreatePrescription(&prescription)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPrescription)
}

func (h *MedicationHandler) GetPrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	prescription, err := h.service.GetPrescriptionByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Prescription not found"})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

func (h *MedicationHandler) DiscontinuePrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req struct {
		Reason string `json:"reason" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prescription, err := h.service.DiscontinuePrescription(uint(id), req.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

func (h *MedicationHandler) ListPatientPrescriptions(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	activeOnly := c.Query("active") == "true"
	var prescriptions []*models.Prescription
	var errSvc error

	if activeOnly {
		prescriptions, errSvc = h.service.ListActivePrescriptions(uint(patientID))
	} else {
		prescriptions, errSvc = h.service.ListPatientPrescriptions(uint(patientID))
	}

	if errSvc != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errSvc.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}
