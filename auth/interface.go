package auth

import "net/http"

type Service interface {
	BasicAuthMiddleware(next http.Handler) http.Handler
}

type Repository interface {
	GetUsernameAndPasswordAdmin(username, password string) (isAdmin bool, err error)
}
