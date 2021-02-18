package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"
	"agent/app/user/service/userclient"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginResp, error) {
	resp, err := l.svcCtx.UserRpc.Login(
		l.ctx, &userclient.LoginReq{
			Phone:    req.Phone,
			Password: req.Password,
			Captcha:  req.Captcha,
		},
	)
	if err != nil {
		return nil, err
	}

	authLogic := NewAuthLogic(l.ctx, l.svcCtx)
	token, err := authLogic.Auth()
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Code:    resp.Code,
		Message: resp.Message,
		Token:   *token,
	}, nil
}
