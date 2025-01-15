package main

import (
	databases "tes/database"
	"tes/database/migration"
	"tes/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//INITIAL DATABASE
	databases.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()

	//manggil nama package
	route.RouteInit(app)

	app.Listen(":8080")
}
