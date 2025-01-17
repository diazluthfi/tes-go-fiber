package middleware

import (
	"tes/utils"

	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware adalah middleware untuk memeriksa apakah token valid
// dan apakah pengguna memiliki peran yang sesuai (dalam hal ini "admin").
func AuthMiddleware(ctx *fiber.Ctx) error {
	// Mengambil token dari header "x-token".
	token := ctx.Get("x-token")

	// Jika token kosong, pengguna dianggap tidak terautentikasi.
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated", // Pesan untuk pengguna bahwa mereka tidak diautentikasi.
		})
	}

	// _, err := utils.VerifToken(token)

	// Memeriksa apakah token valid menggunakan utils.DecodeToken.
	claims, err := utils.DecodeToken(token)
	if err != nil {

		// Jika token tidak valid atau ada kesalahan, kembalikan status 401 Unauthorized.

		if err.Error() == "Token is expired" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired", // Pesan untuk pengguna bahwa token expired.
			})
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated", // Pesan untuk pengguna bahwa token tidak valid.
		})
	}

	// Mengambil klaim "role" dari token untuk memeriksa peran pengguna.
	role := claims["role"].(string)
	// Jika peran pengguna bukan "admin", kembalikan status 403 Forbidden.
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access", // Pesan untuk pengguna bahwa mereka tidak memiliki izin.
		})
	}

	// mengirim data key ke getall
	// ctx.Locals("userInfo", claims)
	// ctx.Locals("role", claims["role"])

	// Jika semua validasi berhasil, middleware melanjutkan ke proses berikutnya.
	return ctx.Next()
}

// PermissionCreate adalah middleware kosong saat ini.
// Fungsinya hanya melanjutkan proses tanpa melakukan validasi tambahan.
func PermissionCreate(ctx *fiber.Ctx) error {
	// Fungsi ini dapat diimplementasikan untuk logika izin di masa depan.
	// Saat ini hanya meneruskan proses ke handler berikutnya.
	return ctx.Next()
}
