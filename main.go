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