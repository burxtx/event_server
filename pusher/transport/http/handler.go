package http

import (
	"context"
	"encoding/json"
	"net/http"

	mylog "git.n.xiaomi.com/op-basic/event_server/libs/log"
	"git.n.xiaomi.com/op-basic/event_server/users"
	"git.n.xiaomi.com/op-basic/event_server/users/auth"
	"github.com/go-chi/chi"
)

type userHandler struct {
	s      auth.Service
	logger mylog.MyLogger
}

func (h *userHandler) router() chi.Router {
	r := chi.NewRouter()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", h.Signup)
	})
	return r
}

func (h *userHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var request struct {
		username string
		password string
		email    string
	}
	u, err := h.s.Register(ctx, request.username, request.password)
	if err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
	var response = struct {
		ID users.UserID `json:"user_id"`
	}{
		ID: u.ID,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.GetLogger(ctx).Errorln(err.Error())
		encodeError(ctx, err, w)
		return
	}
}
