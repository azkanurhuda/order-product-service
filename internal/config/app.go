package config

import (
	"github.com/azkanurhuda/order-product-service/internal/delivery/http"
	"github.com/azkanurhuda/order-product-service/internal/delivery/http/route"
	"github.com/azkanurhuda/order-product-service/internal/repository"
	"github.com/azkanurhuda/order-product-service/internal/usecase/order"
	"github.com/azkanurhuda/order-product-service/sdk/http/costumer_sdk"
	"github.com/azkanurhuda/order-product-service/sdk/http/product_sdk"
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
	// setup sdk
	costumerSDK := costumer_sdk.NewCostumerSDK(config.Config, config.Log)
	productSDK := product_sdk.NewProductSDK(config.Config, config.Log)

	// setup repositories
	repo := repository.NewRepository(config.DB)

	// setup use cases
	orderUseCase := order.NewOrderUseCase(config.DB, config.Log, config.Validate, repo, costumerSDK, productSDK)

	// setup controller
	orderController := http.NewOrderController(config.Log, orderUseCase)

	routeConfig := route.RouteConfig{
		App:             config.App,
		OrderController: orderController,
	}

	routeConfig.Setup()
}
