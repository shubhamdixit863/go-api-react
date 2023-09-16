package dto

import (
	"fmt"
	"goapibackend/internal/domain/entity"
)

type UserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Location  string `json:"location"`
	Schedule  string `json:"schedule"`
	Password  string `json:"password"`
	Degree    string `json:"degree"`
}

type SignInDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userdto *UserDto) ToEntity() *entity.User {
	fmt.Println("Arrived here")
	return &entity.User{
		FirstName:   userdto.FirstName,
		LastName:    userdto.LastName,
		Email:       userdto.Email,
		DegreeLevel: userdto.Degree,
		Password:    userdto.Password,
		Schedule:    userdto.Schedule,
		Location:    userdto.Location,
	}
}

func (userdto *UserDto) FromEntity(userEntity entity.User) {
	userdto.Password = userEntity.Password
	userdto.FirstName = userEntity.FirstName
	userdto.LastName = userEntity.LastName
	userdto.Email = userEntity.Email
	userdto.Degree = userEntity.DegreeLevel
	userdto.Schedule = userEntity.Schedule
	userdto.Location = userEntity.Location

}
