package services

import "goapibackend/internal/domain/dto"

type IUserService interface {
	Signup(userDto *dto.UserDto) (uint, error)
	GetAllUsers() ([]dto.UserDto, error)
}
