package main

import (
	"github.com/gofiber/fiber/v2"
	"golang2/handler" // pastikan ini sesuai dengan struktur folder kamu
)

func main() {
	app := fiber.New()

	// Route GET untuk cek server
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("✅ Server aktif. Gunakan POST /login dan /register.")
	})

	// Route login dan register
	app.Post("/login", handler.LoginHandler)
	app.Post("/register", handler.RegisterHandler)

	// Jalankan server
	app.Listen(":3000")
}
