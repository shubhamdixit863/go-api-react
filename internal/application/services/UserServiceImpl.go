package services

import (
	"fmt"
	"goapibackend/internal/domain/repository"
)

type UserServiceImpl struct {
	// The repository dependency
	UserRepository repository.IUserRepository
}

func (user UserServiceImpl) Signup() {
	user.UserRepository.AddUser()
	fmt.Println("Service MEthod Invoked")

}
