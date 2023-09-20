package dto

import (
	"goapibackend/internal/domain/entity"
	"time"
)

type UserProjectDto struct {
	ProjectName string `json:"projectName"`
	Description string `json:"description"`
	FileName    string `json:"fileName"` // Would be coming from file Service
}

func (Up UserProjectDto) ToEntity() entity.UserProject {

	return entity.UserProject{
		Description: Up.Description,
		ProjectName: Up.ProjectName,
		FileName:    Up.FileName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

}