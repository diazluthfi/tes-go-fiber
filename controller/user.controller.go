package controller

import "github.com/gofiber/fiber/v2"

func userControllerCreate(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"Hello": "World",
	})

}
