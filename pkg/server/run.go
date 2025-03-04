package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() error {
	service := newService()
	mux := http.NewServeMux()
	mux.HandleFunc("/proxy", handleProxy)
	// add middleware to one route
	mux.HandleFunc("/add", middlewareAuthz(service.handleAdd, "admin"))
	// add middleware to all routes
	handler := middlewareLogging(mux)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}

func middlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log request
		slog.Info("request received", "method", r.Method, "path", r.URL.Path)

		// log headers
		for k, v := range r.Header {
			slog.Info("header", k, v)
		}

		next.ServeHTTP(w, r)
	})
}
