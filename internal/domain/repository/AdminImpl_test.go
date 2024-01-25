package repository

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

func DbConnect() *gorm.DB {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalln("Cant find env")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DB")), // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                                                                          // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {

		log.Fatalln("Error in connecting with db", err)
	}

	fmt.Println("Connected with Db")
	return db
}

func TestAdminRepositoryImpl_GetUserById(t *testing.T) {
	repo := AdminRepositoryImpl{Db: DbConnect()}
	user, err := repo.GetUserById(6)
	assert.Nil(t, err)
	assert.Equal(t, "shubham", user.FirstName)

}
