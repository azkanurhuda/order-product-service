package http

import (
	"github.com/azkanurhuda/product-service/internal/model"
	"github.com/azkanurhuda/product-service/internal/usecase/product"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ProductController struct {
	Log     *logrus.Logger
	UseCase product.UseCase
}

func NewProductController(logger *logrus.Logger, useCase product.UseCase) *ProductController {
	return &ProductController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *ProductController) Create(ctx *fiber.Ctx) error {
	request := new(model.ProductRequest)
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

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) ReadByID(ctx *fiber.Ctx) error {
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

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) ReadAll(ctx *fiber.Ctx) error {
	response, err := c.UseCase.ReadByAll(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to register user : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.ProductResponse]{
		Data: response,
	})
}

func (c *ProductController) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		c.Log.Warnf("Failed to convert ID to int: %+v", err)
		return fiber.ErrBadRequest
	}

	request := new(model.ProductRequest)
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

	return ctx.JSON(model.WebResponse[*model.ProductResponse]{Data: response})
}

func (c *ProductController) Delete(ctx *fiber.Ctx) error {
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
