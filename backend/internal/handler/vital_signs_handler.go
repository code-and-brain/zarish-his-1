package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/code-and-brain/zarish-his-1/backend/internal/models"
	"github.com/code-and-brain/zarish-his-1/backend/internal/service"
)

type VitalSignsHandler struct {
	service *service.VitalSignsService
}

func NewVitalSignsHandler(service *service.VitalSignsService) *VitalSignsHandler {
	return &VitalSignsHandler{service: service}
}

func (h *VitalSignsHandler) CreateVitalSigns(c *gin.Context) {
	var vitals models.VitalSigns
	if err := c.ShouldBindJSON(&vitals); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdVitals, err := h.service.CreateVitalSigns(&vitals)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdVitals)
}

func (h *VitalSignsHandler) GetVitalSigns(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	vitals, err := h.service.GetVitalSignsByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vital signs not found"})
		return
	}

	c.JSON(http.StatusOK, vitals)
}

func (h *VitalSignsHandler) ListEncounterVitalSigns(c *gin.Context) {
	encounterID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter ID"})
		return
	}

	vitals, err := h.service.ListEncounterVitalSigns(uint(encounterID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vitals)
}

func (h *VitalSignsHandler) ListPatientVitalSigns(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	vitals, err := h.service.ListPatientVitalSigns(uint(patientID), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vitals)
}
