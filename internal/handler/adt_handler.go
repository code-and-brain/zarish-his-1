package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type ADTHandler struct {
	service *service.ADTService
}

func NewADTHandler(service *service.ADTService) *ADTHandler {
	return &ADTHandler{service: service}
}

func (h *ADTHandler) CreateWard(c *gin.Context) {
	var ward models.Ward
	if err := c.ShouldBindJSON(&ward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateWard(&ward); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ward)
}

func (h *ADTHandler) ListWards(c *gin.Context) {
	wards, err := h.service.ListWards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wards)
}

func (h *ADTHandler) CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, room)
}

func (h *ADTHandler) CreateBed(c *gin.Context) {
	var bed models.Bed
	if err := c.ShouldBindJSON(&bed); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateBed(&bed); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bed)
}

func (h *ADTHandler) ListBeds(c *gin.Context) {
	status := c.Query("status")
	beds, err := h.service.ListBeds(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, beds)
}

func (h *ADTHandler) AdmitPatient(c *gin.Context) {
	var admission models.Admission
	if err := c.ShouldBindJSON(&admission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.AdmitPatient(&admission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, admission)
}

func (h *ADTHandler) DischargePatient(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DischargePatient(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Patient discharged successfully"})
}

func (h *ADTHandler) ListActiveAdmissions(c *gin.Context) {
	admissions, err := h.service.ListActiveAdmissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admissions)
}
