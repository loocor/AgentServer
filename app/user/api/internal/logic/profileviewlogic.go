package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"
	"agent/app/user/service/userclient"
)

type ProfileViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProfileViewLogic {
	return ProfileViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileViewLogic) ProfileView(req types.ProfileViewReq) (*types.CommonResp, error) {
	resp, err := l.svcCtx.UserRpc.ProfileView(
		l.ctx, &userclient.ProfileViewReq{
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
