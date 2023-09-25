package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goapibackend/internal/apis/handlers"
	"goapibackend/internal/apis/middlewares"
	"goapibackend/internal/application/services"
	"goapibackend/internal/domain/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DbConnection() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=berry.db.elephantsql.com user=lwafzeku password=8s3aTiSlqis6ENkJvcZpyZ0yxiAA0XeZ dbname=lwafzeku  sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                                                            // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {

		log.Fatalln("Error in connecting with db", err)
	}

	fmt.Println("Connected with Db")
	return db
}

func main() {

	r := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization"}
	r.Use(cors.New(config))

	//r.Use(cors.Default())
	connectionDb := DbConnection()
	userRepository := repository.UserImpl{Db: connectionDb}
	userService := services.UserServiceImpl{
		UserRepository: &userRepository,
	}
	// Dependency injection
	handler := handlers.Handler{
		UserService: userService,
	}
	r.POST("/signup", handler.SignUp)
	r.POST("/project", handler.AddProject)
	r.GET("/project", middlewares.Authorize(), handler.GetProject)

	r.POST("/signin", handler.SignIn)

	r.GET("/users", handler.GetAllUsers)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
