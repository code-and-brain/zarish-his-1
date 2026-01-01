package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type ClinicalNoteHandler struct {
	service *service.ClinicalNoteService
}

func NewClinicalNoteHandler(service *service.ClinicalNoteService) *ClinicalNoteHandler {
	return &ClinicalNoteHandler{service: service}
}

func (h *ClinicalNoteHandler) CreateNote(c *gin.Context) {
	var note models.ClinicalNote
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdNote, err := h.service.CreateNote(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdNote)
}

func (h *ClinicalNoteHandler) GetNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	note, err := h.service.GetNoteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *ClinicalNoteHandler) UpdateNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	var note models.ClinicalNote
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.ID = uint(id)
	updatedNote, err := h.service.UpdateNote(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedNote)
}

func (h *ClinicalNoteHandler) SignNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	// In a real app, get userID from context/token
	// For MVP, we'll assume a default user or pass it in body
	var signRequest struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&signRequest); err != nil {
		// Default to user 1 if not provided
		signRequest.UserID = 1
	}

	note, err := h.service.SignNote(uint(id), signRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *ClinicalNoteHandler) ListEncounterNotes(c *gin.Context) {
	encounterID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid encounter ID"})
		return
	}

	notes, err := h.service.ListEncounterNotes(uint(encounterID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func (h *ClinicalNoteHandler) ListPatientNotes(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	notes, err := h.service.ListPatientNotes(uint(patientID), limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}
