package entity

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model definition
// Our entity which would be saved in the database

type User struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	FirstName   string
	LastName    string
	Email       string
	DegreeLevel string //bachelors,masters,professor
	Password    string
	Schedule    string
	Location    string
	CreatedAt   time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"default:current_timestamp"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
