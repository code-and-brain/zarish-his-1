package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"zarish-his/backend/internal/domain/models"
	"zarish-his/backend/internal/service/clinical"
)

type DocumentHandler struct {
	service *clinical.DocumentService
}

func NewDocumentHandler(service *clinical.document.DocumentService) *DocumentHandler {
	return &DocumentHandler{service: service}
}

func (h *DocumentHandler) UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	patientID, err := strconv.ParseUint(c.PostForm("patient_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	doc := &models.Document{
		PatientID:   uint(patientID),
		FileName:    file.Filename,
		FileSize:    file.Size,
		ContentType: file.Header.Get("Content-Type"),
	}

	// In a real application, you would save the file to a storage service (e.g., S3)
	// and store the URL in the document model.
	// For this example, we'll just log it.

	createdDoc, err := h.service.SaveDocument(doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDoc)
}

func (h *DocumentHandler) GetDocument(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid document ID"})
		return
	}

	doc, err := h.service.GetDocumentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Document not found"})
		return
	}

	c.JSON(http.StatusOK, doc)
}

func (h *DocumentHandler) ListPatientDocuments(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	docs, err := h.service.ListPatientDocuments(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, docs)
}
