package logic

import (
	"context"

	"agent/app/user/service/internal/svc"
	"agent/app/user/service/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.CommonResp, error) {
	return &user.CommonResp{
		Code:    0,
		Message: "logout touched",
	}, nil
}
