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
			Id:        int(data.ID),
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

func (ad *AdminServiceImpl) GetUserById(id int) (dto.UserDto, error) {
	// We call the repo
	data, err := ad.AdminRepo.GetUserById(id)
	if err != nil {
		return dto.UserDto{}, err
	}
	d := dto.UserDto{
		Id:        int(data.ID),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Location:  data.Location,
		Schedule:  data.Schedule,
		Password:  data.Password,
		Degree:    data.DegreeLevel,
	}
	return d, nil
}
