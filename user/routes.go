package user

import (
	"github.com/go-chi/chi"
	"kriya_test_backend/auth"
)

type Routes struct {
	delivery *HTTPDelivery
	auth     auth.Service
}

func NewRoutes(delivery *HTTPDelivery, auth auth.Service) *Routes {
	return &Routes{delivery: delivery, auth: auth}
}

func (routes Routes) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(routes.auth.BasicAuthMiddleware)
			r.Route("/users", func(r chi.Router) {
				r.Post("/create", routes.delivery.PostUser)
				r.Put("/update", routes.delivery.UpdateUser)
				r.Delete("/delete", routes.delivery.DeleteUser)
			})
		})
		
		r.Route("/user", func(r chi.Router) {
			r.Get("/list", routes.delivery.GetListUser)
			r.Get("/detail", routes.delivery.GetUser)
		})
	})
}
