package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/code-and-brain/zarish-his-1/backend/internal/models"
	"github.com/code-and-brain/zarish-his-1/backend/internal/service"
)

type LabHandler struct {
	service *service.LabService
}

func NewLabHandler(service *service.LabService) *LabHandler {
	return &LabHandler{service: service}
}

func (h *LabHandler) CreateLabTest(c *gin.Context) {
	var test models.LabTest
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTest, err := h.service.CreateLabTest(&test)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTest)
}

func (h *LabHandler) ListLabTests(c *gin.Context) {
	tests, err := h.service.ListLabTests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tests)
}

func (h *LabHandler) CreateLabOrder(c *gin.Context) {
	var order models.LabOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdOrder, err := h.service.CreateLabOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdOrder)
}

func (h *LabHandler) GetLabOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	order, err := h.service.GetLabOrderByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lab order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *LabHandler) ListPatientLabOrders(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	orders, err := h.service.ListPatientLabOrders(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *LabHandler) AddLabResult(c *gin.Context) {
	var result models.LabResult
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdResult, err := h.service.AddLabResult(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdResult)
}
