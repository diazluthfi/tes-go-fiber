package route

import (
	"tes/config"
	"tes/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Get("/user", handlers.UserhandlersRead)
	r.Get("/user/:id", handlers.GetHandlerById)
	r.Post("/usercreate", handlers.UserHandlerCreate)
	r.Put("/userupdate/:id", handlers.UserHandlerUpdate)
	r.Put("/userupdate/:id/email", handlers.UserHandlerUpdateEmail)
	r.Delete("/userdelete/:id", handlers.UserHandlerDelete)

}
