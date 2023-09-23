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
	Projects     []UserProjectDto `json:"projects"`
	TotalRecords int64            `json:"totalRecords"`
	TotalPages   int64            `json:"totalPages"`
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
