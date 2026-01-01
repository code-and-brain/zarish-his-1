package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string   `json:"username" gorm:"uniqueIndex;not null"`
	Password  string   `json:"-" gorm:"not null"` // Hashed password
	Email     string   `json:"email" gorm:"uniqueIndex"`
	Role      string   `json:"role" gorm:"default:'PATIENT'"` // ADMIN, DOCTOR, NURSE, PATIENT
	PatientID *uint    `json:"patient_id"`
	Patient   *Patient `json:"patient" gorm:"foreignKey:PatientID"`
	Active    bool     `json:"active" gorm:"default:true"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
