package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"goapibackend/internal/domain/dto"
	"goapibackend/internal/domain/repository"
	"math"
	"time"
)

type UserServiceImpl struct {
	// The repository dependency
	UserRepository repository.IUserRepository
}

func (user UserServiceImpl) GetAllProjects(page, limit int) (*dto.UserProjectDtoResponse, error) {

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
	totalPages := int64(math.Ceil(float64(count) / float64(limit)))

	return &dto.UserProjectDtoResponse{
		Projects:     userProjects,
		TotalRecords: count,
		TotalPages:   totalPages,
	}, nil

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

// Method to create a jwt here ----https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-New-Hmac

func CreateJwt(userName string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userName,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":      time.Now().Add(10 * time.Minute).UnixNano(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (user UserServiceImpl) SignIn(userDto *dto.SignInDto) (string, error) {

	returnedUser, err := user.UserRepository.GetUserByEmail(userDto.Email)
	if err != nil {
		return "", err
	}
	// We will check for password too
	if returnedUser.Password != userDto.Password {
		return "", errors.New("UserName Or Pasword Dono't Match")
	}
	// we have to issue the token here ---
	createJwt, err := CreateJwt(userDto.Email)
	if err != nil {
		return "", err
	}
	return createJwt, nil
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
