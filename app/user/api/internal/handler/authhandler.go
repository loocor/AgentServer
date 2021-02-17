package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"

	"agent/app/user/api/internal/logic"
	"agent/app/user/api/internal/svc"
)

func AuthHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewAuthLogic(r.Context(), ctx)
		resp, err := l.Auth()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
