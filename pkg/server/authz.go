package server

import (
	"log/slog"
	"net/http"
)

func authzmiddleware(next http.HandlerFunc, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRoles := getRoles(r)
		for _, role := range roles {
			for _, userRole := range userRoles {
				if role == userRole {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		slog.Debug("forbidden", "roles", userRoles)
		http.Error(w, "forbidden", http.StatusForbidden)
	})
}

func getRoles(r *http.Request) []string {
	return r.Header["X-Roles"]
}
