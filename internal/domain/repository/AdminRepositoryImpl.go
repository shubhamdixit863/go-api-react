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
	tx := ad.Db.Find(&users).Limit(limit).Offset(page * limit)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil

}
