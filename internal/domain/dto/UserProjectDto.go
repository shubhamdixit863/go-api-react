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

type UserProjectDtoResponse struct {
	ProjectName string `json:"projectName"`
	Description string `json:"description"`
	FileName    string `json:"fileName"` // Would be coming from file Service
	Count       int    `json:"count"`
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
