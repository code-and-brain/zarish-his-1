package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/code-and-brain/zarish-his-1/backend/internal/service"
)

type ReportingHandler struct {
	service *service.ReportingService
}

func NewReportingHandler(service *service.ReportingService) *ReportingHandler {
	return &ReportingHandler{service: service}
}

func (h *ReportingHandler) GetDailyOPDReport(c *gin.Context) {
	dateStr := c.Query("date")
	var date time.Time
	var err error

	if dateStr == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
	}

	report, err := h.service.GenerateDailyOPDReport(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *ReportingHandler) GetDiseaseSurveillanceReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	now := time.Now()
	startDate := now.AddDate(0, 0, -30) // Default last 30 days
	endDate := now

	var err error
	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format. Use YYYY-MM-DD"})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format. Use YYYY-MM-DD"})
			return
		}
	}

	report, err := h.service.GenerateDiseaseSurveillanceReport(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
