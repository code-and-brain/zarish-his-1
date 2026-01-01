package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zarishsphere/zarish-his/internal/handler"
	"github.com/zarishsphere/zarish-his/internal/models"
	"github.com/zarishsphere/zarish-his/internal/repository"
	"github.com/zarishsphere/zarish-his/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Read configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "zarish_his"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate models
	log.Println("Running database migrations...")
	err = db.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Encounter{},
		&models.VitalSigns{},
		&models.ClinicalNote{},
		&models.Invoice{},
		&models.InvoiceItem{},
		&models.Payment{},
		&models.InsuranceClaim{},
		&models.ImagingStudy{},
		&models.ImagingSeries{},
		&models.ImagingInstance{},
		&models.RadiologyReport{},
		&models.Medication{},
		&models.Prescription{},
		&models.LabTest{},
		&models.LabOrder{},
		&models.LabResult{},
		&models.Appointment{},
		&models.Ward{},
		&models.Room{},
		&models.Bed{},
		&models.Admission{},
		&models.Transfer{},
		&models.DischargeSummary{},
		&models.PharmacyStock{},
		&models.Dispensing{},
		&models.StockMovement{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Repositories
	patientRepo := repository.NewPatientRepository(db)
	encounterRepo := repository.NewEncounterRepository(db)
	vitalSignsRepo := repository.NewVitalSignsRepository(db)
	clinicalNoteRepo := repository.NewClinicalNoteRepository(db)
	medicationRepo := repository.NewMedicationRepository(db)
	labRepo := repository.NewLabRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)

	// Initialize Services
	patientService := service.NewPatientService(patientRepo)
	encounterService := service.NewEncounterService(encounterRepo)
	vitalSignsService := service.NewVitalSignsService(vitalSignsRepo)
	clinicalNoteService := service.NewClinicalNoteService(clinicalNoteRepo)

	cdsService := service.NewCDSService()
	medicationService := service.NewMedicationService(medicationRepo, cdsService)

	labService := service.NewLabService(labRepo)
	appointmentService := service.NewAppointmentService(appointmentRepo)

	// Initialize Handlers
	patientHandler := handler.NewPatientHandler(patientService)
	encounterHandler := handler.NewEncounterHandler(encounterService)
	vitalSignsHandler := handler.NewVitalSignsHandler(vitalSignsService)
	clinicalNoteHandler := handler.NewClinicalNoteHandler(clinicalNoteService)
	medicationHandler := handler.NewMedicationHandler(medicationService, patientService)
	labHandler := handler.NewLabHandler(labService)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	// Initialize ADT
	adtRepo := repository.NewADTRepository(db)
	adtService := service.NewADTService(adtRepo)
	adtHandler := handler.NewADTHandler(adtService)

	billingRepo := repository.NewBillingRepository(db)
	billingService := service.NewBillingService(billingRepo)
	billingHandler := handler.NewBillingHandler(billingService)

	radiologyRepo := repository.NewRadiologyRepository(db)
	radiologyService := service.NewRadiologyService(radiologyRepo)
	radiologyHandler := handler.NewRadiologyHandler(radiologyService)

	// Initialize Reporting
	reportingService := service.NewReportingService(db)
	reportingHandler := handler.NewReportingHandler(reportingService)

	// Initialize Pharmacy
	pharmacyRepo := repository.NewPharmacyRepository(db)
	pharmacyService := service.NewPharmacyService(pharmacyRepo)
	pharmacyHandler := handler.NewPharmacyHandler(pharmacyService)

	// Initialize Portal
	portalHandler := handler.NewPortalHandler(
		patientService,
		appointmentService,
		labService,
		medicationService,
		clinicalNoteService,
	)

	// Setup Router
	r := gin.Default()

	// CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := r.Group("/api/v1")
	{
		// Insurance Routes
		insurance := api.Group("/insurance")
		{
			insurance.POST("/claims", billingHandler.SubmitClaim)
			insurance.GET("/claims/:id", billingHandler.GetClaim)
			insurance.GET("/claims/pending", billingHandler.GetPendingClaims)
			insurance.POST("/claims/:id/approve", billingHandler.ApproveClaim)
			insurance.POST("/claims/:id/reject", billingHandler.RejectClaim)
		}

		// Radiology Routes
		radiology := api.Group("/radiology")
		{
			radiology.POST("/studies", radiologyHandler.CreateStudy)
			radiology.GET("/studies", radiologyHandler.ListStudies)
			radiology.GET("/studies/:id", radiologyHandler.GetStudy)
			radiology.PUT("/studies/:id/status", radiologyHandler.UpdateStatus)
			radiology.GET("/worklist", radiologyHandler.GetWorklist)

			radiology.POST("/reports", radiologyHandler.CreateReport)
			radiology.PUT("/reports/:id", radiologyHandler.UpdateReport)
		}
		// Patient Routes
		api.POST("/patients", patientHandler.CreatePatient)
		api.GET("/patients/:id", patientHandler.GetPatient)
		api.PUT("/patients/:id", patientHandler.UpdatePatient)
		api.DELETE("/patients/:id", patientHandler.DeletePatient)
		api.GET("/patients", patientHandler.ListPatients)
		api.GET("/patients/search", patientHandler.SearchPatients)
		api.GET("/patients/:id/history", patientHandler.GetPatientHistory)

		// Encounter Routes
		api.POST("/encounters", encounterHandler.CreateEncounter)
		api.GET("/encounters/:id", encounterHandler.GetEncounter)
		api.PUT("/encounters/:id", encounterHandler.UpdateEncounter)
		api.PUT("/encounters/:id/status", encounterHandler.UpdateStatus)
		api.GET("/patients/:id/encounters", encounterHandler.ListPatientEncounters)

		// Vital Signs Routes
		api.POST("/vital-signs", vitalSignsHandler.CreateVitalSigns)
		api.GET("/vital-signs/:id", vitalSignsHandler.GetVitalSigns)
		api.GET("/encounters/:id/vital-signs", vitalSignsHandler.ListEncounterVitalSigns)
		api.GET("/patients/:id/vital-signs", vitalSignsHandler.ListPatientVitalSigns)

		// Clinical Notes Routes
		api.POST("/clinical-notes", clinicalNoteHandler.CreateNote)
		api.GET("/clinical-notes/:id", clinicalNoteHandler.GetNote)
		api.PUT("/clinical-notes/:id", clinicalNoteHandler.UpdateNote)
		api.POST("/clinical-notes/:id/sign", clinicalNoteHandler.SignNote)
		api.GET("/encounters/:id/clinical-notes", clinicalNoteHandler.ListEncounterNotes)
		api.GET("/patients/:id/clinical-notes", clinicalNoteHandler.ListPatientNotes)

		// Medication Routes
		api.POST("/medications", medicationHandler.CreateMedication)
		api.GET("/medications/search", medicationHandler.SearchMedications)

		// Prescription Routes
		api.POST("/prescriptions", medicationHandler.CreatePrescription)
		api.GET("/prescriptions/:id", medicationHandler.GetPrescription)
		api.POST("/prescriptions/:id/discontinue", medicationHandler.DiscontinuePrescription)
		api.GET("/patients/:id/prescriptions", medicationHandler.ListPatientPrescriptions)

		// Lab Routes
		api.POST("/lab-tests", labHandler.CreateLabTest)
		api.GET("/lab-tests", labHandler.ListLabTests)
		api.POST("/lab-orders", labHandler.CreateLabOrder)
		api.GET("/lab-orders/:id", labHandler.GetLabOrder)
		api.GET("/patients/:id/lab-orders", labHandler.ListPatientLabOrders)
		api.POST("/lab-results", labHandler.AddLabResult)

		// Appointment Routes
		api.POST("/appointments", appointmentHandler.CreateAppointment)
		api.GET("/appointments/:id", appointmentHandler.GetAppointment)
		api.PUT("/appointments/:id", appointmentHandler.UpdateAppointment)
		api.POST("/appointments/:id/cancel", appointmentHandler.CancelAppointment)
		api.GET("/appointments", appointmentHandler.ListAppointments)
		api.GET("/patients/:id/appointments", appointmentHandler.ListPatientAppointments)

		// ADT Routes
		api.POST("/wards", adtHandler.CreateWard)
		api.GET("/wards", adtHandler.ListWards)
		api.POST("/rooms", adtHandler.CreateRoom)
		api.POST("/beds", adtHandler.CreateBed)
		api.GET("/beds", adtHandler.ListBeds)
		api.POST("/admissions", adtHandler.AdmitPatient)
		api.POST("/admissions/:id/discharge", adtHandler.DischargePatient)
		api.GET("/admissions/active", adtHandler.ListActiveAdmissions)
		api.GET("/admissions/:id", adtHandler.GetAdmission)
		api.POST("/transfers", adtHandler.TransferPatient)
		api.GET("/transfers", adtHandler.ListTransfers)
		api.POST("/discharge-summaries", adtHandler.CreateDischargeSummary)
		api.GET("/admissions/:id/discharge-summary", adtHandler.GetDischargeSummary)

		// Reporting Routes
		api.GET("/reports/daily-opd", reportingHandler.GetDailyOPDReport)
		api.GET("/reports/disease-surveillance", reportingHandler.GetDiseaseSurveillanceReport)

		// Pharmacy Routes
		api.POST("/pharmacy/stock", pharmacyHandler.AddStock)
		api.GET("/pharmacy/stock/:medication_id", pharmacyHandler.GetStock)
		api.GET("/pharmacy/stock/low", pharmacyHandler.GetLowStock)
		api.POST("/pharmacy/dispense", pharmacyHandler.DispenseMedication)
		api.GET("/pharmacy/dispensing-queue", pharmacyHandler.GetDispensingQueue)
		api.GET("/pharmacy/history/:patient_id", pharmacyHandler.GetPatientHistory)
		api.GET("/pharmacy/movements/:medication_id", pharmacyHandler.GetStockMovements)

		// Portal Routes
		portal := api.Group("/portal")
		{
			portal.GET("/dashboard", portalHandler.GetDashboard)
			portal.GET("/appointments", portalHandler.GetAppointments)
			portal.GET("/records", portalHandler.GetRecords)
		}
	}

	log.Printf("Zarish-HIS server starting on :%s", port)
	r.Run(":" + port)
}
