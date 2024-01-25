package services

import "goapibackend/internal/domain/dto"

type IAdminService interface {
	GetUser(page, limit int) ([]dto.UserDto, error)
	GetUserById(id int) (dto.UserDto, error)
}
