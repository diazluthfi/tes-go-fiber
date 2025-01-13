package main

import (
	"tes/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//manggil nama package
	route.RouteInit(app)

	app.Listen(":8080")
}
