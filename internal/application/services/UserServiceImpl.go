package services

import (
	"errors"
	"fmt"
	"goapibackend/internal/domain/dto"
	"goapibackend/internal/domain/repository"
)

type UserServiceImpl struct {
	// The repository dependency
	UserRepository repository.IUserRepository
}

func (user UserServiceImpl) GetAllProjects(page, limit int) ([]dto.UserProjectDto, error) {

	// Call the service method
	projects, count, err := user.UserRepository.GetProjects(page, limit)
	fmt.Println(count)
	if err != nil {
		return nil, err
	}
	var userProjects []dto.UserProjectDto
	for _, v := range projects {
		up := dto.UserProjectDto{
			ProjectName: v.ProjectName,
			Description: v.Description,
			FileName:    v.FileName,
		}
		userProjects = append(userProjects, up)
	}
	return userProjects, nil

}

func (user UserServiceImpl) AddProject(projectDto *dto.UserProjectDto) (uint, error) {
	projectId, err := user.UserRepository.AddProject(projectDto.ToEntity())
	if err != nil {
		return 0, err
	}

	return projectId, nil
}

func (user UserServiceImpl) Signup(userDto *dto.UserDto) (uint, error) {

	// We will convert here
	addUser, err := user.UserRepository.AddUser(userDto.ToEntity())
	if err != nil {
		return addUser, err
	}
	return addUser, nil

}

func (user UserServiceImpl) SignIn(userDto *dto.SignInDto) (uint, error) {

	returnedUser, err := user.UserRepository.GetUserByEmail(userDto.Email)
	if err != nil {
		return 0, err
	}
	// We will check for password too
	if returnedUser.Password != userDto.Password {
		return 0, errors.New("UserName Or Pasword Dono't Match")
	}
	return 0, nil
}

func (user UserServiceImpl) GetAllUsers() ([]dto.UserDto, error) {
	// return the data by calling repository
	var userDtoAll []dto.UserDto
	users, err := user.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(users); i++ {
		var userdto dto.UserDto
		userdto.FromEntity(users[i])
		userDtoAll = append(userDtoAll, userdto)
	}
	return userDtoAll, nil
}
