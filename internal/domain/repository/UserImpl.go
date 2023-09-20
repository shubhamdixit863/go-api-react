package repository

import (
	"errors"
	"goapibackend/internal/domain/entity"
	"gorm.io/gorm"
)

type UserImpl struct {
	// Will have our database connection
	Db *gorm.DB
}

func (ui *UserImpl) AddProject(userProject entity.UserProject) (uint, error) {

	// Our Db query To save the User Project inside the db
	tx := ui.Db.Create(&userProject)
	if tx.Error != nil {
		return userProject.ID, tx.Error
	}
	return userProject.ID, nil

}

func (ui *UserImpl) AddUser(user *entity.User) (uint, error) {
	var alreadyExisting entity.User
	// We will check here first if the user with email already exists
	tx := ui.Db.Where("email = ?", user.Email).First(&alreadyExisting)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		// We wil save in our db
		tx = ui.Db.Create(user)
		if tx.Error != nil {
			return user.ID, tx.Error
		}

		return user.ID, nil
	}

	return user.ID, errors.New("User With Email Already Exists")

}

func (ui *UserImpl) GetAllUsers() ([]entity.User, error) {
	var users []entity.User
	tx := ui.Db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ui *UserImpl) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	tx := ui.Db.Where("email = ?", email).First(&user)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (ui *UserImpl) AutoMigrate() error {
	// We will write the migration part
	userEntity := entity.User{}
	userProjectEntity := entity.UserProject{}
	err := ui.Db.AutoMigrate(&userEntity, &userProjectEntity)
	if err != nil {
		return err
	}

	return nil
}
