package service

import (
	"golang2/model"
)

var users []model.User

// Register user baru
func Register(user model.User) (bool, string) {
	// Cek jika username sudah ada
	for _, u := range users {
		if u.Username == user.Username {
			return false, "Username sudah digunakan"
		}
	}
	users = append(users, user)
	return true, "Registrasi berhasil"
}

// Login user
func Login(user model.User) (bool, string) {
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			return true, "Login berhasil"
		}
	}
	return false, "Username atau password salah"
}

