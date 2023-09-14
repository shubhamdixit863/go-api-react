package services

import (
	"goapibackend/internal/domain/dto"
	"goapibackend/internal/domain/repository"
)

type UserServiceImpl struct {
	// The repository dependency
	UserRepository repository.IUserRepository
}

func (user UserServiceImpl) Signup(userDto *dto.UserDto) (uint, error) {

	// We will convert here
	addUser, err := user.UserRepository.AddUser(userDto.ToEntity())
	if err != nil {
		return addUser, err
	}
	return addUser, nil

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
