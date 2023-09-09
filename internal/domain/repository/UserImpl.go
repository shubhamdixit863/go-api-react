package repository

import (
	"fmt"
	"goapibackend/internal/domain/entity"
	"gorm.io/gorm"
)

type UserImpl struct {
	// Will have our database connection
	Db *gorm.DB
}

func (ui *UserImpl) AddUser() {
	fmt.Println("User Repository invoked")

	// Logic For Db operation
}

func (ui *UserImpl) AutoMigrate() error {
	// We will write the migration part
	userEntity := entity.User{}
	err := ui.Db.AutoMigrate(&userEntity)
	if err != nil {
		return err
	}

	return nil
}
