package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type PharmacyHandler struct {
	service *service.PharmacyService
}

func NewPharmacyHandler(service *service.PharmacyService) *PharmacyHandler {
	return &PharmacyHandler{service: service}
}

func (h *PharmacyHandler) AddStock(c *gin.Context) {
	var stock models.PharmacyStock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddStock(&stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

func (h *PharmacyHandler) GetStock(c *gin.Context) {
	medicationID, _ := strconv.Atoi(c.Param("medication_id"))

	stocks, err := h.service.GetAvailableStock(uint(medicationID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func (h *PharmacyHandler) GetLowStock(c *gin.Context) {
	stocks, err := h.service.GetLowStockAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func (h *PharmacyHandler) DispenseMedication(c *gin.Context) {
	var dispensing models.Dispensing
	if err := c.ShouldBindJSON(&dispensing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DispenseMedication(&dispensing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dispensing)
}

func (h *PharmacyHandler) GetDispensingQueue(c *gin.Context) {
	prescriptions, err := h.service.GetDispensingQueue()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}

func (h *PharmacyHandler) GetPatientHistory(c *gin.Context) {
	patientID, _ := strconv.Atoi(c.Param("patient_id"))

	history, err := h.service.GetPatientDispensingHistory(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

func (h *PharmacyHandler) GetStockMovements(c *gin.Context) {
	medicationID, _ := strconv.Atoi(c.Param("medication_id"))

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
			return
		}
	}

	movements, err := h.service.GetStockMovementReport(uint(medicationID), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movements)
}
