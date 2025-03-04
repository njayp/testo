package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleWare(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("x-test", "hello")

	w := httptest.NewRecorder()
	handler := middlewareLogging(http.NotFoundHandler())
	handler.ServeHTTP(w, req)
}
