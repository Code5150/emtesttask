//go:build wireinject
// +build wireinject

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/google/wire"

	"gorm.io/gorm"

	"emtesttask/config"
	"emtesttask/controller"
	"emtesttask/mapper"
	"emtesttask/repository"
	"emtesttask/service"
)

func provideAppConfig() (*config.AppConfig, error) {
	return config.LoadConfig(".")
}

func provideDatabaseConfig(cfg *config.AppConfig) (*config.DatabaseConfig, error) {
	dbConfig := new(config.DatabaseConfig{
		DSN:          fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort),
		MaxOpenConns: 17,
		MaxIdleConns: 10,
	})
	return dbConfig, nil
}

func provideDatabase(cfg *config.DatabaseConfig) (*config.Database, func(), error) {
	return config.NewDatabase(cfg)
}

func provideGormDB(db *config.Database) *gorm.DB {
	return db.DB
}

func provideRouter(subscriptionController *controller.SubscriptionController) *gin.Engine {
	router := gin.Default()
	subscriptionController.RegisterRoutes(router)
	return router
}

func InitializeApp() (*gin.Engine, func(), error) {
	wire.Build(
		// База данных
		provideAppConfig,
		provideDatabaseConfig,
		provideDatabase,
		provideGormDB,

		// Репозиторий
		repository.NewSubscriptionRepository,

		// Маппер
		mapper.NewSubscriptionMapper,

		// Сервис
		service.NewSubscriptionService,

		// Контроллер
		controller.NewSubscriptionController,

		// Роутер
		provideRouter,
	)

	return nil, nil, nil
}
