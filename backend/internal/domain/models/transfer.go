package models

import "time"

// Transfer represents an internal patient transfer between wards/beds
type Transfer struct {
	BaseModel

	AdmissionID uint       `gorm:"not null;index" json:"admission_id"`
	Admission   *Admission `gorm:"foreignKey:AdmissionID" json:"admission,omitempty"`

	FromWardID uint  `gorm:"not null" json:"from_ward_id"`
	FromWard   *Ward `gorm:"foreignKey:FromWardID" json:"from_ward,omitempty"`

	FromBedID uint `gorm:"not null" json:"from_bed_id"`
	FromBed   *Bed `gorm:"foreignKey:FromBedID" json:"from_bed,omitempty"`

	ToWardID uint  `gorm:"not null" json:"to_ward_id"`
	ToWard   *Ward `gorm:"foreignKey:ToWardID" json:"to_ward,omitempty"`

	ToBedID uint `gorm:"not null" json:"to_bed_id"`
	ToBed   *Bed `gorm:"foreignKey:ToBedID" json:"to_bed,omitempty"`

	TransferDate time.Time `gorm:"not null" json:"transfer_date"`
	Reason       string    `gorm:"type:text" json:"reason"`
	AuthorizedBy uint      `gorm:"not null" json:"authorized_by"` // UserID
}

func (Transfer) TableName() string {
	return "transfers"
}
