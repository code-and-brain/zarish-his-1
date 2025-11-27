package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type EncounterHandler struct {
	service *service.EncounterService
}

func NewEncounterHandler(service *service.EncounterService) *EncounterHandler {
	return &EncounterHandler{service: service}
}

func (h *EncounterHandler) CreateEncounter(c *gin.Context) {
	var encounter models.Encounter
	if err := c.ShouldBindJSON(&encounter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdEncounter, err := h.service.CreateEncounter(&encounter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdEncounter)
}

func (h *EncounterHandler) GetEncounter(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter ID"})
		return
	}

	encounter, err := h.service.GetEncounterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Encounter not found"})
		return
	}

	c.JSON(http.StatusOK, encounter)
}

func (h *EncounterHandler) UpdateEncounter(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter ID"})
		return
	}

	var encounter models.Encounter
	if err := c.ShouldBindJSON(&encounter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encounter.ID = uint(id)
	updatedEncounter, err := h.service.UpdateEncounter(&encounter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedEncounter)
}

func (h *EncounterHandler) ListPatientEncounters(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	encounters, total, err := h.service.ListPatientEncounters(uint(patientID), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  encounters,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *EncounterHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter ID"})
		return
	}

	var statusUpdate struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var encounter *models.Encounter
	if statusUpdate.Status == "in-progress" {
		encounter, err = h.service.StartEncounter(uint(id))
	} else if statusUpdate.Status == "finished" {
		encounter, err = h.service.FinishEncounter(uint(id))
	} else {
		// Generic update for other statuses
		encounter, err = h.service.GetEncounterByID(uint(id))
		if err == nil {
			encounter.Status = statusUpdate.Status
			encounter, err = h.service.UpdateEncounter(encounter)
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, encounter)
}
