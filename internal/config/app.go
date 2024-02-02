package config

import (
	"github.com/azkanurhuda/costumer-service/internal/delivery/http"
	"github.com/azkanurhuda/costumer-service/internal/delivery/http/route"
	"github.com/azkanurhuda/costumer-service/internal/repository"
	"github.com/azkanurhuda/costumer-service/internal/usecase/costumer"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	repo := repository.NewRepository(config.DB)

	// setup use cases
	costumerUseCase := costumer.NewCostumerUseCase(config.DB, config.Log, config.Validate, repo)

	// setup controller
	costumerController := http.NewCostumerController(config.Log, costumerUseCase)

	routeConfig := route.RouteConfig{
		App:                config.App,
		CostumerController: costumerController,
	}

	routeConfig.Setup()
}
