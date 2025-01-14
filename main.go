package main

import (
	"tes/database"
	"tes/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//INITIAL DATABASE
	database.DatabaseInit()
	app := fiber.New()

	//manggil nama package
	route.RouteInit(app)

	app.Listen(":8080")
}
