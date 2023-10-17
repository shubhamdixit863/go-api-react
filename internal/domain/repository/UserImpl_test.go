package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func DbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=goapibackend-db.cake0vsnlrsi.us-east-1.rds.amazonaws.com user=postgres password=Q3j2eynfz8mlUioFzReH dbname=postgres  sslmode=require TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                                                                // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
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
