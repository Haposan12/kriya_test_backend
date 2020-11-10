package user

import "github.com/go-chi/chi"

type Routes struct {
	delivery *HTTPDelivery
}

func NewRoutes(delivery *HTTPDelivery) *Routes {
	return &Routes{delivery: delivery}
}

func (routes Routes) RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/list", routes.delivery.GetListUser)
			r.Get("/detail", routes.delivery.GetUser)
		})
	})
}
