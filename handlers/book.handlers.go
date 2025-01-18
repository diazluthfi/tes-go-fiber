package handlers

import (
	"fmt"
	"log"
	databases "tes/database"
	"tes/model/entity"
	"tes/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {

	book := new(request.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
	}

	//handle file
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("error file", errFile)
	}

	var filename string
	if file != nil {
		filename = file.Filename

		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))
		if errSaveFile != nil {
			log.Println("save file failed")
		}
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filename,
	}

	errCreateBook := databases.DB.Debug().Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"Message": "Failed to Store",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}
