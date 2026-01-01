package models

import (
	"time"

	"gorm.io/gorm"
)

// ImagingStudy represents a DICOM Study or FHIR ImagingStudy resource
type ImagingStudy struct {
	gorm.Model
	StudyUID          string     `json:"study_uid" gorm:"uniqueIndex;not null"`
	AccessionNumber   string     `json:"accession_number" gorm:"uniqueIndex"`
	PatientID         uint       `json:"patient_id" gorm:"index"`
	Patient           Patient    `json:"patient" gorm:"foreignKey:PatientID"`
	EncounterID       *uint      `json:"encounter_id" gorm:"index"`
	Encounter         *Encounter `json:"encounter" gorm:"foreignKey:EncounterID"`
	Modality          string     `json:"modality"` // e.g., CT, MRI, DX
	BodySite          string     `json:"body_site"`
	Description       string     `json:"description"`
	Status            string     `json:"status"` // scheduled, in-progress, completed, cancelled
	StartedAt         time.Time  `json:"started_at"`
	CompletedAt       *time.Time `json:"completed_at"`
	NumberOfSeries    int        `json:"number_of_series"`
	NumberOfInstances int        `json:"number_of_instances"`
	ReferrerID        *uint      `json:"referrer_id"`        // Practitioner ID
	PerformingTechID  *uint      `json:"performing_tech_id"` // Practitioner ID

	Series []ImagingSeries  `json:"series" gorm:"foreignKey:StudyID"`
	Report *RadiologyReport `json:"report" gorm:"foreignKey:StudyID"`
}

// ImagingSeries represents a DICOM Series
type ImagingSeries struct {
	gorm.Model
	SeriesUID   string            `json:"series_uid" gorm:"uniqueIndex;not null"`
	StudyID     uint              `json:"study_id" gorm:"index"`
	Number      int               `json:"number"`
	Modality    string            `json:"modality"`
	Description string            `json:"description"`
	BodyPart    string            `json:"body_part"`
	Instances   []ImagingInstance `json:"instances" gorm:"foreignKey:SeriesID"`
}

// ImagingInstance represents a DICOM Instance (Image)
type ImagingInstance struct {
	gorm.Model
	SOPInstanceUID string `json:"sop_instance_uid" gorm:"uniqueIndex;not null"`
	SeriesID       uint   `json:"series_id" gorm:"index"`
	Number         int    `json:"number"`
	Title          string `json:"title"`
	FilePath       string `json:"file_path"`    // Path to file storage or PACS URL
	ContentType    string `json:"content_type"` // e.g., application/dicom, image/jpeg
}

// RadiologyReport represents the findings for a study
type RadiologyReport struct {
	gorm.Model
	StudyID       uint       `json:"study_id" gorm:"uniqueIndex"`
	RadiologistID uint       `json:"radiologist_id"` // Practitioner ID
	Status        string     `json:"status"`         // draft, preliminary, final, amended
	Findings      string     `json:"findings" gorm:"type:text"`
	Impression    string     `json:"impression" gorm:"type:text"`
	Conclusion    string     `json:"conclusion" gorm:"type:text"`
	ReportedAt    time.Time  `json:"reported_at"`
	FinalizedAt   *time.Time `json:"finalized_at"`
}
