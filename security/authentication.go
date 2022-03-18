package security

import (
	"context"
	"log"
	"net/http"

	"github.com/giancarlobastos/loteca-backend/service"
)

type AuthenticationMiddleware struct {
	apiService *service.ApiService
}

func NewAuthenticationMiddleware(as *service.ApiService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		apiService: as,
	}
}

func (amw *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("[AuthenticationMiddleware] - panic occurred:", err)
				http.NotFound(w, r)
			}
		}()

		token := r.Header.Get("Token")
		user, err := amw.apiService.Authenticate(token)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", *user)))
	})
}
