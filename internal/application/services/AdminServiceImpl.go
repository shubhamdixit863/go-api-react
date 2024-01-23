package services

import (
	"goapibackend/internal/domain/dto"
	"goapibackend/internal/domain/repository"
)

type AdminServiceImpl struct {
	AdminRepo repository.AdminRepository
}

func (ad *AdminServiceImpl) GetUser(page, limit int) ([]dto.UserDto, error) {
	var users []dto.UserDto
	user, err := ad.AdminRepo.GetUser(page, limit)
	if err != nil {
		return nil, err
	}
	for _, data := range user {
		users = append(users, dto.UserDto{
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
			Location:  data.Location,
			Schedule:  data.Schedule,
			Password:  data.Password,
			Degree:    data.DegreeLevel,
		})

	}
	return users, nil
}
