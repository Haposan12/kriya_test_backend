package main

import (
	"github.com/go-chi/chi"
	"kriya_test_backend/auth"
	"kriya_test_backend/config"
	"kriya_test_backend/user"
	"kriya_test_backend/utils"
)

func main() {
	authRepository := auth.NewRepository(config.GlobalDBSQL)
	authService := auth.NewService(*authRepository)

	userRepository := user.NewRepository(config.GlobalDBSQL)
	userService := user.NewService(*userRepository)
	userController := user.NewDelivery(*userService)
	userRoutes := user.NewRoutes(userController, authService)

	r := chi.NewRouter()

	userRoutes.RegisterRoutes(r)

	port := ":12345"

	utils.Listen(port, r)
}
