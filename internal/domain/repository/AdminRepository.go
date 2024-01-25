package repository

import (
	"goapibackend/internal/domain/entity"
)

type AdminRepository interface {
	GetUser(page, limit int) ([]entity.User, error)
	GetUserById(id int) (entity.User, error)
}
