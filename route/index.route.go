package route

import (
	"tes/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", handlers.UserhandlersRead)
	r.Post("/usercreate", handlers.UserHandlerCreate)

}
