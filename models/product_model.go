package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint     `gorm:"primary_key"`
	Title       string   `gorm:"not null"`
	Description string   `gorm:"not null"`
	Images      []string `gorm:"type:text[]"`
	Notes       string
	Latitude    float64  `gorm:"not null"`
	Longitude   float64  `gorm:"not null"`
	Metadata    []string `gorm:"type:text[]"`
	CollectedBy string
	Likes       int
	Views       int
	Flags       int
	Validity    time.Time `gorm:"not null"`
	PostedOn    time.Time `gorm:"not null"`
	PostedBy    string    `gorm:"not null"`
	ShowOnMap   bool      `gorm:"not null"`
	Archive     []string  `gorm:"type:text[]"`
	Status      string    `gorm:"not null"`
}
