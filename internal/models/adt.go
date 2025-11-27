package models

import (
	"time"

	"gorm.io/gorm"
)

type Ward struct {
	gorm.Model
	Name        string `json:"name"`
	Department  string `json:"department"`
	Type        string `json:"type"` // e.g., General, ICU, Private
	Description string `json:"description"`
	Rooms       []Room `json:"rooms,omitempty"`
}

type Room struct {
	gorm.Model
	WardID      uint   `json:"ward_id"`
	RoomNumber  string `json:"room_number"`
	Type        string `json:"type"` // e.g., Single, Double, Shared
	Description string `json:"description"`
	Beds        []Bed  `json:"beds,omitempty"`
}

type Bed struct {
	gorm.Model
	RoomID    uint   `json:"room_id"`
	BedNumber string `json:"bed_number"`
	Status    string `json:"status" gorm:"default:'Available'"` // Available, Occupied, Maintenance, Cleaning
	Type      string `json:"type"`                              // e.g., Standard, ICU, Pediatric
	Notes     string `json:"notes"`
}

type Admission struct {
	gorm.Model
	PatientID         uint      `json:"patient_id"`
	Patient           Patient   `json:"patient,omitempty"`
	BedID             uint      `json:"bed_id"`
	Bed               Bed       `json:"bed,omitempty"`
	AdmissionDate     time.Time `json:"admission_date"`
	DischargeDate     *time.Time `json:"discharge_date,omitempty"`
	AdmittingDoctorID uint      `json:"admitting_doctor_id"`
	Diagnosis         string    `json:"diagnosis"`
	Status            string    `json:"status" gorm:"default:'Admitted'"` // Admitted, Discharged, Transferred
	Notes             string    `json:"notes"`
}
