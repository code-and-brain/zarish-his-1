package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type RadiologyHandler struct {
	service *service.RadiologyService
}

func NewRadiologyHandler(service *service.RadiologyService) *RadiologyHandler {
	return &RadiologyHandler{service: service}
}

func (h *RadiologyHandler) CreateStudy(c *gin.Context) {
	var study models.ImagingStudy
	if err := c.ShouldBindJSON(&study); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateStudy(&study); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, study)
}

func (h *RadiologyHandler) GetStudy(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	study, err := h.service.GetStudy(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Study not found"})
		return
	}

	c.JSON(http.StatusOK, study)
}

func (h *RadiologyHandler) ListStudies(c *gin.Context) {
	patientID, _ := strconv.Atoi(c.Query("patient_id"))
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	studies, total, err := h.service.ListStudies(uint(patientID), status, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  studies,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *RadiologyHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Status string `json:"status"`
		TechID uint   `json:"tech_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if req.Status == "in-progress" {
		err = h.service.StartExam(uint(id), req.TechID)
	} else if req.Status == "completed" {
		err = h.service.CompleteExam(uint(id))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status transition"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

func (h *RadiologyHandler) CreateReport(c *gin.Context) {
	var report models.RadiologyReport
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateReport(&report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, report)
}

func (h *RadiologyHandler) UpdateReport(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var report models.RadiologyReport
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report.ID = uint(id)
	if err := h.service.UpdateReport(&report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *RadiologyHandler) GetWorklist(c *gin.Context) {
	studies, err := h.service.GetWorklist()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, studies)
}
