package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthz(t *testing.T) {
	for _, test := range []struct {
		name  string
		roles []string
		code  int
	}{
		{
			name:  "not authorized",
			roles: []string{"racoon"},
			code:  http.StatusForbidden,
		},
		{
			name:  "authorized",
			roles: []string{"admin"},
			code:  http.StatusOK,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-Roles", test.roles[0])

			w := httptest.NewRecorder()
			handler := middlewareAuthz(func(w http.ResponseWriter, r *http.Request) {}, "admin")
			handler.ServeHTTP(w, req)

			if w.Code != test.code {
				t.Errorf("expected %d, got %d", test.code, w.Code)
			}
		})
	}
}
