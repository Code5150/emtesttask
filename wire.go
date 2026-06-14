//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"

	"github.com/google/wire"

	"gorm.io/gorm"

	"emtesttask/config"
	"emtesttask/controller"
	"emtesttask/mapper"
	"emtesttask/repository"
	"emtesttask/service"
)

func provideDatabaseConfig() (*config.DatabaseConfig, error) {
	config := new(config.DatabaseConfig{
		DSN:          "host=localhost user=postgres password=Us3r1pa55w0rd dbname=ef_mob port=5432 sslmode=disable",
		MaxOpenConns: 17,
		MaxIdleConns: 10,
	})
	return config, nil
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
