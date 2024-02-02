package route

import (
	"github.com/azkanurhuda/product-service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App               *fiber.App
	ProductController *http.ProductController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/products", c.ProductController.Create)
	c.App.Get("/api/product/:id", c.ProductController.ReadByID)
	c.App.Get("/api/products", c.ProductController.ReadAll)
	c.App.Put("/api/product/:id", c.ProductController.Update)
	c.App.Delete("/api/product/:id", c.ProductController.Delete)
}
