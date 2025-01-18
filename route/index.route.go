package route

import (
	"tes/config"
	"tes/handlers"
	"tes/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Post("/login", handlers.AuthHandlersLogin)
	r.Get("/user", middleware.AuthMiddleware, handlers.UserhandlersRead)
	r.Get("/user/:id", handlers.GetHandlerById)
	r.Post("/usercreate", handlers.UserHandlerCreate)
	r.Put("/userupdate/:id", handlers.UserHandlerUpdate)
	r.Put("/userupdate/:id/email", handlers.UserHandlerUpdateEmail)
	r.Delete("/userdelete/:id", handlers.UserHandlerDelete)
	r.Post("/book", handlers.BookHandlerCreate)
}
