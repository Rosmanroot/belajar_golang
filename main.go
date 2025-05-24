package main

import "github.com/gofiber/fiber/v2"

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func ProsesLogin(input LoginInput) (bool, string) {
	// Username dan password yang dianggap benar
	if input.Username == "admin" && input.Password == "123" {
		return true, "Login berhasil"
	}
	return false, "Username atau password salah"
}

func main() {
	app := fiber.New()

	// Route GET "/" (untuk akses dari browser)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ Server aktif. Gunakan POST /login untuk login.")
	})

	// Route POST "/login" (untuk proses login)
	app.Post("/login", func(c *fiber.Ctx) error {
		var input LoginInput

		// Parsing body JSON ke struct
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Format JSON tidak valid",
			})
		}

		// Panggil fungsi login
		success, msg := ProsesLogin(input)

		if success {
			return c.JSON(fiber.Map{
				"message": msg,
				"user":    input.Username,
			})
		}
		return c.Status(401).JSON(fiber.Map{
			"message": msg,
		})
	})

	// Jalankan server di port 3000
	app.Listen(":3000")
}
