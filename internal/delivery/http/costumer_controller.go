package http

import (
	"github.com/azkanurhuda/costumer-service/internal/model"
	"github.com/azkanurhuda/costumer-service/internal/usecase/costumer"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

type CostumerController struct {
	Log     *logrus.Logger
	UseCase costumer.UseCase
}

func NewCostumerController(logger *logrus.Logger, useCase costumer.UseCase) *CostumerController {
	return &CostumerController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CostumerController) Create(ctx *fiber.Ctx) error {
	request := new(model.CostumerRequest)
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

	return ctx.JSON(model.WebResponse[*model.CostumerResponse]{Data: response})
}

func (c *CostumerController) ReadByID(ctx *fiber.Ctx) error {
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

	return ctx.JSON(model.WebResponse[*model.CostumerResponse]{Data: response})
}

func (c *CostumerController) ReadAll(ctx *fiber.Ctx) error {
	response, err := c.UseCase.ReadByAll(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CostumerResponse]{
		Data: response,
	})
}

func (c *CostumerController) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Log.Warnf("Failed to convert ID to int: %+v", err)
		return fiber.ErrBadRequest
	}

	request := new(model.CostumerRequest)
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

	return ctx.JSON(model.WebResponse[*model.CostumerResponse]{Data: response})
}

func (c *CostumerController) Delete(ctx *fiber.Ctx) error {
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
