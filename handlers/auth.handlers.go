package handlers

import (
	"log"                    // Untuk logging error
	databases "tes/database" // Mengimport koneksi database dari folder `database`
	"tes/model/entity"       // Mengimport model entity, seperti `User`
	"tes/model/request"      // Mengimport struktur data untuk request, seperti `LoginRequest`
	"tes/utils"              // Mengimport utilitas, seperti hashing dan token
	"time"                   // Untuk manipulasi waktu

	"github.com/dgrijalva/jwt-go"            // Untuk pembuatan dan verifikasi JWT
	"github.com/go-playground/validator/v10" // Untuk validasi input
	"github.com/gofiber/fiber/v2"            // Framework web untuk Go
)

// Fungsi untuk menangani login
func AuthHandlersLogin(ctx *fiber.Ctx) error {
	// Menginisialisasi struktur LoginRequest untuk membaca data input dari body
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		// Mengembalikan error jika parsing body gagal
		return err
	}

	// Validasi password menggunakan library validator
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		// Mengembalikan error jika validasi input gagal
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
	}

	// Mencari user di database berdasarkan email
	var user entity.User
	err := databases.DB.First(&user, "email=?", loginRequest.Email).Error
	if err != nil {
		// Jika email tidak ditemukan, mengembalikan status Unauthorized
		log.Printf("Login failed for email: %s, invalid credentials.", loginRequest.Email)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email",
		})

	}

	// Mengecek apakah password yang dimasukkan sesuai dengan password yang di-hash di database
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		// Jika password salah, mengembalikan status Unauthorized
		log.Printf("Login failed for email: %s, invalid credentials.", loginRequest.Email)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password.",
		})

	}

	// Membuat klaim untuk token JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name                             // Menyimpan nama user
	claims["email"] = user.Email                           // Menyimpan email user
	claims["address"] = user.Address                       // Menyimpan alamat user
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix() // Menyimpan waktu kedaluwarsa token (2 menit dari sekarang)

	// Menentukan role berdasarkan email
	if user.Email == "te2113a@gmail.com" {
		claims["role"] = "admin" // Jika email adalah milik admin, role diatur sebagai admin
	} else {
		claims["role"] = "user" // Selain itu, role diatur sebagai user
	}

	// Membuat token JWT menggunakan klaim di atas
	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		// Jika token gagal dibuat, mencatat error dan mengembalikan status Unauthorized
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	// Mengembalikan token JWT sebagai respon
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
