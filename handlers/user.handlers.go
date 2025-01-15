package handlers

import (
	"log"
	databases "tes/database"
	"tes/model/entity"
	"tes/model/request"

	"github.com/gofiber/fiber/v2"
)

func UserhandlersRead(ctx *fiber.Ctx) error {

	var user []entity.User
	result := databases.DB.Debug().Find(&user)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(user)

}

func UserHandlerCreate(ctx *fiber.Ctx) error {

	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// validate := validator.New()
	// errValidate := validate.Struct(user)
	// if errValidate != nil {
	// 	return ctx.Status(400).JSON(fiber.Map{
	// 		"message": "Failed",
	// 		"error":   errValidate.Error(),
	// 	})
	// }

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	errCreateUser := databases.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message": "Failed to Store",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})

}
