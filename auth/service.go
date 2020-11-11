package auth

import (
	"kriya_test_backend/shared/response"
	"net/http"
)

type ServiceImpl struct {
	authRepository Repository
}

func NewService(authRepository Repository) *ServiceImpl {
	return &ServiceImpl{authRepository: authRepository}
}

func (svc ServiceImpl) BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			realm := "Please enter your username and password for this site"
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			response.SendErrorResponse(w, http.StatusText(http.StatusUnauthorized), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		var err error
		isAdmin, err := svc.authRepository.GetUsernameAndPasswordAdmin(username, password)

		if err != nil {
			response.SendErrorResponse(w, err.Error(), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if !isAdmin {
			response.SendErrorResponse(w, "you dont have access", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// Middleware logic ends here
		next.ServeHTTP(w, r)
	})
}
