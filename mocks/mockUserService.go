package mocks

import (
	"github.com/stretchr/testify/mock"

	"goapibackend/internal/domain/dto"
)

type MockUserService struct {
	mock.Mock
}

func (m MockUserService) GetAllProjects(page, limit int) (*dto.UserProjectDtoResponse, error) {
	m.Called(page, limit)
	return nil, nil
}

func (m MockUserService) AddProject(projectDto *dto.UserProjectDto) (uint, error) {
	m.Called(projectDto)
	return 0, nil
}

func (m MockUserService) Signup(userDto *dto.UserDto) (uint, error) {
	args := m.Called(userDto)
	return args.Get(0).(uint), args.Error(1)
}

func (m MockUserService) SignIn(userDto *dto.SignInDto) (string, error) {
	m.Called(userDto)
	return "", nil
}

func (m MockUserService) GetAllUsers() ([]dto.UserDto, error) {
	args := m.Called()
	return args.Get(0).([]dto.UserDto), args.Error(1)
}
