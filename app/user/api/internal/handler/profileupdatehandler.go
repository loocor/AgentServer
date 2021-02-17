package handler

import (
	"net/http"

	"agent/app/user/api/internal/logic"
	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ProfileUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProfileUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProfileUpdateLogic(r.Context(), ctx)
		resp, err := l.ProfileUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
