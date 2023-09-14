package repository

import "goapibackend/internal/domain/entity"

type IUserRepository interface {
	AddUser(user *entity.User) (uint, error)
	GetAllUsers() ([]entity.User, error)
	AutoMigrate() error
}
