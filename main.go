package main

import (
	"github.com/go-chi/chi"
	"kriya_test_backend/config"
	"kriya_test_backend/user"
	"kriya_test_backend/utils"
)

func main() {
	repository := user.NewRepository(config.GlobalDBSQL)

	service := user.NewService(*repository)

	delivery := user.NewDelivery(*service)

	routes := user.NewRoutes(delivery)

	r := chi.NewRouter()

	routes.RegisterRoutes(r)

	port := ":12345"

	utils.Listen(port, r)
}
