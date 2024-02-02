package http

import (
	"github.com/azkanurhuda/order-product-service/internal/model"
	"github.com/azkanurhuda/order-product-service/internal/usecase/order"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

type OrderController struct {
	Log     *logrus.Logger
	UseCase order.UseCase
}

func NewOrderController(logger *logrus.Logger, useCase order.UseCase) *OrderController {
	return &OrderController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *OrderController) Create(ctx *fiber.Ctx) error {
	request := new(model.OrderRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OrderResponse]{Data: response})
}

func (c *OrderController) ReadByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Log.Warnf("Failed to convert ID to int: %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.ReadByID(ctx.UserContext(), ID)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OrderFullResponse]{Data: response})
}

func (c *OrderController) ReadAll(ctx *fiber.Ctx) error {
	response, err := c.UseCase.ReadByAll(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.OrderResponse]{
		Data: response,
	})
}

func (c *OrderController) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Log.Warnf("Failed to convert ID to int: %+v", err)
		return fiber.ErrBadRequest
	}

	request := new(model.OrderRequest)
	err = ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Update(ctx.UserContext(), request, ID)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.OrderResponse]{Data: response})
}

func (c *OrderController) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Log.Warnf("Failed to convert ID to int: %+v", err)
		return fiber.ErrBadRequest
	}

	err = c.UseCase.Delete(ctx.UserContext(), ID)
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
