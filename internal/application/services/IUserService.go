package services

import "goapibackend/internal/domain/dto"

type IUserService interface {
	Signup(userDto *dto.UserDto) (uint, error)
	SignIn(userDto *dto.SignInDto) (string, error)
	AddProject(projectDto *dto.UserProjectDto) (uint, error)
	GetAllProjects(page, limit int) (*dto.UserProjectDtoResponse, error)
	GetAllUsers() ([]dto.UserDto, error)
}
