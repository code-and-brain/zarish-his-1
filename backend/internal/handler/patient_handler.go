package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zarish-his/backend/internal/models"
	"zarish-his/backend/internal/service"
)

type PatientHandler struct {
	service *service.PatientService
}

func NewPatientHandler(service *service.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

// CreatePatient creates a new patient
// @Summary Create a new patient
// @Description Create a new patient with nationality-based validation
// @Tags patients
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient object"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/patients [post]
func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate based on nationality
	if err := patient.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPatient, err := h.service.CreatePatient(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPatient)
}

// GetPatient retrieves a patient by ID
// @Summary Get a patient by ID
// @Description Get a patient by ID
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Failure 404 {object} map[string]string
// @Router /api/v1/patients/{id} [get]
func (h *PatientHandler) GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	patient, err := h.service.GetPatientByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient updates a patient
// @Summary Update a patient
// @Description Update a patient by ID
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Patient object"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/patients/{id} [put]
func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient.ID = uint(id)

	// Validate based on nationality
	if err := patient.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPatient, err := h.service.UpdatePatient(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPatient)
}

// ListPatients lists all patients with pagination and filtering
// @Summary List patients
// @Description List all patients with optional filters
// @Tags patients
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Param nationality query string false "Filter by nationality"
// @Param search query string false "Search by name, MRN, phone"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/patients [get]
func (h *PatientHandler) ListPatients(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	nationality := c.Query("nationality")
	search := c.Query("search")

	patients, total, err := h.service.ListPatients(page, limit, nationality, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  patients,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// SearchPatients searches patients by various criteria
// @Summary Search patients
// @Description Search patients by name, MRN, NID, UNHCR number, phone
// @Tags patients
// @Produce json
// @Param q query string true "Search query"
// @Success 200 {array} models.Patient
// @Router /api/v1/patients/search [get]
func (h *PatientHandler) SearchPatients(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	patients, err := h.service.SearchPatients(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

// GetPatientHistory gets complete patient history
// @Summary Get patient history
// @Description Get complete patient history including encounters, medications, labs
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /api/v1/patients/{id}/history [get]
func (h *PatientHandler) GetPatientHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	history, err := h.service.GetPatientHistory(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

// DeletePatient soft deletes a patient
// @Summary Delete a patient
// @Description Soft delete a patient by ID
// @Tags patients
// @Param id path int true "Patient ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /api/v1/patients/{id} [delete]
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	if err := h.service.DeletePatient(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
