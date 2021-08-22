package http

import (
	"context"
	"encoding/json"
	"net/http"

	mylog "git.n.xiaomi.com/op-basic/event_server/libs/log"
	"git.n.xiaomi.com/op-basic/event_server/users/auth"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTPServer struct {
	userSvc auth.Service

	logger mylog.MyLogger

	router chi.Router
}

func NewHTTPServer(userSvc auth.Service, logger mylog.MyLogger) *HTTPServer {
	s := &HTTPServer{}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Route("/user", func(r chi.Router) {
		h := userHandler{s.userSvc, logger}
		r.Mount("/v1", h.router())
	})

	r.Method("GET", "/metrics", promhttp.Handler())
	s.router = r
	return s
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
