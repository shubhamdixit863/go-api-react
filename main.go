package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"goapibackend/internal/apis/handlers"
	"goapibackend/internal/apis/middlewares"
	"goapibackend/internal/application/services"
	"goapibackend/internal/domain/repository"
)

func DbConnection() *gorm.DB {
	fmt.Println(fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DB")))
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

func main() {

	godotenv.Load()

	//r.Use(cors.Default())
	connectionDb := DbConnection()
	r := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

	// Admin routes
	adminHandler := handlers.AdminHandler{
		AdminService: &services.AdminServiceImpl{AdminRepo: &repository.AdminRepositoryImpl{Db: connectionDb}},
	}

	admin := r.Group("/admin")
	{
		admin.GET("/users", adminHandler.GetUsers)
		admin.GET("/user/:id", adminHandler.GetUser)

	}

	userRepository := repository.UserImpl{Db: connectionDb}
	err := userRepository.AutoMigrate()
	if err != nil {
		log.Fatalln("error doing migration", err)
	}
	userService := services.UserServiceImpl{
		UserRepository: &userRepository,
	}
	// Dependency injection
	handler := handlers.Handler{
		UserService: userService,
	}
	r.GET("/", handler.Healthcheck)
	r.POST("/signup", handler.SignUp)
	r.POST("/project", handler.AddProject)
	r.GET("/project", middlewares.Authorize(), handler.GetProject)

	r.POST("/signin", handler.SignIn)

	r.GET("/users", handler.GetAllUsers)

	// Admin apis

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
