package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"go-hexagonal-auth/api"
	bniController "go-hexagonal-auth/api/v1/bni"
	userController "go-hexagonal-auth/api/v1/user"
	bniService "go-hexagonal-auth/business/bni"
	userService "go-hexagonal-auth/business/user"
	"go-hexagonal-auth/config"
	bniRepository "go-hexagonal-auth/modules/bni"
	billingRepository "go-hexagonal-auth/modules/billing"
	migration "go-hexagonal-auth/modules/migration"
	clientRepository "go-hexagonal-auth/modules/sangu_bni"
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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgresConnection(cfg *config.Config) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name,
	)
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.InitMigrate(db)
	fmt.Println("Migration Complete")
	return db
}

func newDatabaseConnection(cfg *config.Config) *gorm.DB {

	configDB := map[string]string{
		"DB_Username": cfg.DBSecondary.Username,
		"DB_Password": cfg.DBSecondary.Password,
		"DB_Port":     cfg.DBSecondary.Port,
		"DB_Host":     cfg.DBSecondary.Host,
		"DB_Name":     cfg.DBSecondary.Name,
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

	//load config if available or set to default
	config := config.InitConfig()
	//initialize database connection based on given config
	//dbConnection := newDatabaseConnection(&config)
	dbSecondConnection := newPostgresConnection(&config)


	//initiate user repository
	userRepo := userRepository.NewGormDBRepository(dbSecondConnection)

	//initiate user service
	userService := userService.NewService(userRepo)

	//initiate user controller
	userController := userController.NewController(userService)



	//initiate auth service
	authService := authService.NewService(userService, userRepo, config)

	//initiate auth controller
	authController := authController.NewController(authService, config)

	//initiate auth service
	mediaService := mediaService.NewService(config)

	//initiate media
	mediaController := mediaController.NewController(mediaService, config)

	clientRepo := clientRepository.NewClient()
	clientRepo.ClientID = config.BNIConfig.ClientID
	clientRepo.BaseURL	= config.BNIConfig.Url
	clientRepo.ClientSecret = config.BNIConfig.Key
	//initiate user repository
	bniRepo := bniRepository.NewBNIConfiguration(&config.BNIConfig,dbSecondConnection, &clientRepo)

	//initiate user repository
	billingRepo := billingRepository.NewGormDBRepository(dbSecondConnection, config)

	//initiate user service
	bniService := bniService.NewService(&config,bniRepo, billingRepo)

	//initiate user controller
	bniController := bniController.NewController(bniService, config)

	//create echo http
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{ "http://localhost:3030"},
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

