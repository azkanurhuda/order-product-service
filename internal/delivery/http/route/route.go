package route

import (
	"github.com/azkanurhuda/order-product-service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App             *fiber.App
	OrderController *http.OrderController
	AuthMiddleware  fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/orders", c.OrderController.Create)
	c.App.Get("/api/order/:id", c.OrderController.ReadByID)
	c.App.Get("/api/orders", c.OrderController.ReadAll)
	c.App.Put("/api/order/:id", c.OrderController.Update)
	c.App.Delete("/api/order/:id", c.OrderController.Delete)
}
