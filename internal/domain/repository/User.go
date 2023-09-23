package repository

import "goapibackend/internal/domain/entity"

type IUserRepository interface {
	AddUser(user *entity.User) (uint, error)
	GetAllUsers() ([]entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	AddProject(userProject entity.UserProject) (uint, error)
	GetProjects(currentPage, limit int) ([]entity.UserProject, int64, error)
	AutoMigrate() error
}
