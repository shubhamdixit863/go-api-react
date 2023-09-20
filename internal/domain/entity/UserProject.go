package entity

import "time"

type UserProject struct {
	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	ProjectName string
	Description string
	FileName    string
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`
}
