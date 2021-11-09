package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"go-hexagonal-auth/api"
	bniController "go-hexagonal-auth/api/v1/bni"
	userController "go-hexagonal-auth/api/v1/user"
	adminService "go-hexagonal-auth/business/admin"
	bniService "go-hexagonal-auth/business/bni"
	userService "go-hexagonal-auth/business/user"
	"go-hexagonal-auth/config"
	adminRepository "go-hexagonal-auth/modules/admin"
	bniRepository "go-hexagonal-auth/modules/bni"
	migration "go-hexagonal-auth/modules/migration"
	userRepository "go-hexagonal-auth/modules/user"

	authController "go-hexagonal-auth/api/v1/auth"
	authService "go-hexagonal-auth/business/auth"

	mediaController "go-hexagonal-auth/api/v1/media"
	mediaService "go-hexagonal-auth/business/media"

	"os"
	"os/signal"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabaseConnection(cfg *config.Config) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": cfg.DB.Username,
		"DB_Password": cfg.DB.Password,
		"DB_Port":     cfg.DB.Port,
		"DB_Host":     cfg.DB.Host,
		"DB_Name":     cfg.DB.Name,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	migration.InitMigrate(db)

	return db
}

func main() {
	fmt.Println("OK")
	//load config if available or set to default
	config := config.InitConfig()
	//initialize database connection based on given config
	dbConnection := newDatabaseConnection(&config)



	//initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate user controller
	userController := userController.NewController(userService)

	//initiate user repository
	adminRepo := adminRepository.NewGormDBRepository(dbConnection)

	//initiate user service
	adminService := adminService.NewService(adminRepo)


	//initiate auth service
	authService := authService.NewService(userService, adminService, adminRepo, userRepo, config)

	//initiate auth controller
	authController := authController.NewController(authService, config)

	//initiate auth service
	mediaService := mediaService.NewService(config)

	//initiate media
	mediaController := mediaController.NewController(mediaService, config)

	//initiate user repository
	bniRepo := bniRepository.NewBNIConfiguration(&config.BNIConfig,dbConnection)

	//initiate user service
	bniService := bniService.NewService(&config,bniRepo)

	//initiate user controller
	bniController := bniController.NewController(bniService, config)

	fmt.Println(bniRepo)
	//create echo http
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:9090"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//timeoutContext := time.Duration(config.Server.WriteTimeout) * time.Second


	//register API path and handler
	api.RegisterPath(e, authController, userController,mediaController, bniController, config)

	// run server
	go func() {
		address := fmt.Sprintf("%s", config.Server.Addr)

		e.Static("auth/","public")
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

