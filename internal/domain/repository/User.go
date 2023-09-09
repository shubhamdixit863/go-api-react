package repository

type IUserRepository interface {
	AddUser()
	AutoMigrate() error
}
