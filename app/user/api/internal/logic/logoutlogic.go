package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"
	"agent/app/user/service/userclient"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) LogoutLogic {
	return LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req types.LogoutReq) (*types.CommonResp, error) {
	resp, err := l.svcCtx.UserRpc.Logout(
		l.ctx, &userclient.LogoutReq{
			Phone: req.Phone,
		},
	)

	if err != nil {
		return nil, err
	}

	return &types.CommonResp{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    resp.Data,
	}, nil
}
