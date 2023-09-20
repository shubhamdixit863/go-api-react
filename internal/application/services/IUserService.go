package services

import "goapibackend/internal/domain/dto"

type IUserService interface {
	Signup(userDto *dto.UserDto) (uint, error)
	SignIn(userDto *dto.SignInDto) (uint, error)
	AddProject(projectDto *dto.UserProjectDto) (uint, error)

	GetAllUsers() ([]dto.UserDto, error)
}
