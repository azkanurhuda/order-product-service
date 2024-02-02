package route

import (
	"github.com/azkanurhuda/costumer-service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                *fiber.App
	CostumerController *http.CostumerController
	AuthMiddleware     fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/costumers", c.CostumerController.Create)
	c.App.Get("/api/costumer/:id", c.CostumerController.ReadByID)
	c.App.Get("/api/costumers", c.CostumerController.ReadAll)
	c.App.Put("/api/costumer/:id", c.CostumerController.Update)
	c.App.Delete("/api/costumer/:id", c.CostumerController.Delete)
}
