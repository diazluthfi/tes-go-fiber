package handlers

import (
	"log"
	databases "tes/database"
	"tes/model/entity"
	"tes/model/request"
	"tes/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthHandlersLogin(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	//Validation Passowrd

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User

	err := databases.DB.First(&user, "email=?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Coredential",
		})
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Coredential",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["name"] = user.Email
	claims["name"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Coredential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}
