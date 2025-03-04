package server

import (
	"log/slog"
	"net/http"
	"strings"
)

func Run() error {
	s := newService()
	mux := http.NewServeMux()
	mux.Handle("/add", authzmiddleware(s.HandleAdd, "admin"))
	handle := loggingmiddleware(mux)
	return http.ListenAndServe(":8080", handle)
}

func loggingmiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log request
		slog.Info("request received", "method", r.Method, "path", r.URL.Path)

		// log headers
		for k, v := range r.Header {
			slog.Info("header", k, strings.Join(v, ","))
		}

		next.ServeHTTP(w, r)
	})
}
