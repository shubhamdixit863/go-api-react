package entity

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
// Our entity which would be saved in the database

type User struct {
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	DegreeLevel string //bachelors,masters,professor
	Schedule    string
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
