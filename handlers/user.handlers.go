package handlers

import (
	"log"
	databases "tes/database"
	"tes/model/entity"
	"tes/model/request"
	"tes/model/response"

	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Address:  user.Address,
		Phone:    user.Phone,
		Password: user.Password,
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

func GetHandlerById(ctx *fiber.Ctx) error {

	userId := ctx.Params("id")

	var user entity.User

	err := databases.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	userResponse := response.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    userResponse,
	})

}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
	var user entity.User

	//VALIDATE
	userId := ctx.Params("id")
	err := databases.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	//LOGIKA UPDATE
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	if userRequest.Address != "" {
		user.Address = userRequest.Address
	}
	if userRequest.Phone != "" {
		user.Phone = userRequest.Phone
	}

	errUpdate := databases.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    user,
	})

}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
	var user entity.User
	var isEmailUserExist entity.User
	//VALIDATE
	userId := ctx.Params("id")
	err := databases.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	// Check Email
	errCheckEmail := databases.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error
	if errCheckEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "Email Already Use",
		})
	}

	//LOGIKA UPDATE

	user.Email = userRequest.Email

	errUpdate := databases.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "succes",
		"data":    user,
	})

}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id") // id menyesuaikan dengan yang ada di route
	var user entity.User

	//Check User
	err := databases.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	errDelete := databases.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "User Was Deleted",
	})
}
