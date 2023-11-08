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

func DbConnection() *gorm.DB {
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

func TestDbConnection(t *testing.T) {
	connection := DbConnection()
	assert.NotNil(t, connection)

}

func TestAutoMigrate(t *testing.T) {
	connection := DbConnection()
	repo := UserImpl{Db: connection}
	err := repo.AutoMigrate()
	assert.Nil(t, err)

}
