package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/service"
)

type BillingHandler struct {
	service *service.BillingService
}

func NewBillingHandler(service *service.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

// Invoice Handlers
func (h *BillingHandler) CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.GenerateInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func (h *BillingHandler) GetInvoice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	invoice, err := h.service.GetInvoice(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (h *BillingHandler) GetPatientInvoices(c *gin.Context) {
	patientID, _ := strconv.Atoi(c.Param("patient_id"))

	invoices, err := h.service.GetPatientInvoices(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

func (h *BillingHandler) GetOutstandingInvoices(c *gin.Context) {
	invoices, err := h.service.GetOutstandingInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

// Payment Handlers
func (h *BillingHandler) RecordPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RecordPayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func (h *BillingHandler) GetInvoicePayments(c *gin.Context) {
	invoiceID, _ := strconv.Atoi(c.Param("invoice_id"))

	payments, err := h.service.GetInvoicePayments(uint(invoiceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *BillingHandler) GetPaymentReport(c *gin.Context) {
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
	} else {
		startDate = time.Now().AddDate(0, 0, -30) // Default last 30 days
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
			return
		}
	} else {
		endDate = time.Now()
	}

	payments, err := h.service.GetPaymentReport(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// Insurance Claim Handlers
func (h *BillingHandler) SubmitClaim(c *gin.Context) {
	var claim models.InsuranceClaim
	if err := c.ShouldBindJSON(&claim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SubmitInsuranceClaim(&claim); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, claim)
}

func (h *BillingHandler) GetClaim(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	claim, err := h.service.GetClaim(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Claim not found"})
		return
	}

	c.JSON(http.StatusOK, claim)
}

func (h *BillingHandler) GetPendingClaims(c *gin.Context) {
	claims, err := h.service.GetPendingClaims()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, claims)
}

func (h *BillingHandler) ApproveClaim(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		ApprovedAmount float64 `json:"approved_amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.ApproveClaim(uint(id), req.ApprovedAmount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim approved"})
}

func (h *BillingHandler) RejectClaim(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Reason string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RejectClaim(uint(id), req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Claim rejected"})
}
