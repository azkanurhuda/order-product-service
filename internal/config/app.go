package config

import (
	"github.com/azkanurhuda/product-service/internal/delivery/http"
	"github.com/azkanurhuda/product-service/internal/delivery/http/route"
	"github.com/azkanurhuda/product-service/internal/repository"
	"github.com/azkanurhuda/product-service/internal/usecase/product"
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
	productUseCase := product.NewProductUseCase(config.DB, config.Log, config.Validate, repo)

	// setup controller
	productController := http.NewProductController(config.Log, productUseCase)

	routeConfig := route.RouteConfig{
		App:               config.App,
		ProductController: productController,
	}

	routeConfig.Setup()
}
