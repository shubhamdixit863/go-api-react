package repository

import (
	"gorm.io/gorm"
	"log"

	"goapibackend/internal/domain/entity"
)

type AdminRepositoryImpl struct {
	Db *gorm.DB
}

func (ad *AdminRepositoryImpl) GetUser(page, limit int) ([]entity.User, error) {
	log.Println(page, limit)
	var users []entity.User
	tx := ad.Db.Find(&users).Offset(0).Limit(2)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}

func (ad *AdminRepositoryImpl) GetUserById(id int) (entity.User, error) {
	user := entity.User{
		ID: uint(id),
	}
	result := ad.Db.First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
