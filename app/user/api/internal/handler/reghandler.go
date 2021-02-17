package handler

import (
	"net/http"

	"agent/app/user/api/internal/logic"
	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func RegHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRegLogic(r.Context(), ctx)
		resp, err := l.Reg(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
