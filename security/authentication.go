package security

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/giancarlobastos/loteca-backend/service"
	"github.com/gorilla/mux"
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

		isManagementEndpoint, err := amw.isManagementEndpoint(r)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		token := r.Header.Get("Token")

		if isManagementEndpoint && amw.apiService.AuthenticateManager(token) == nil {
			next.ServeHTTP(w, r)
			return
		} else if isManagementEndpoint {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		user, err := amw.apiService.Authenticate(token)

		if err != nil {
			http.Error(w, "unauthorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", *user)))
	})
}

func (amw *AuthenticationMiddleware) isManagementEndpoint(r *http.Request) (bool, error) {
	route := mux.CurrentRoute(r)
	path, err := route.GetPathTemplate()

	if err != nil {
		log.Printf("Error [isManagementEndpoint] - %v", err)
		return false, err
	}

	return strings.HasPrefix(path, "/manager/"), nil
}
