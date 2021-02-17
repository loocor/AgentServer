package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"
	"agent/app/user/service/userclient"
)

type RegLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegLogic {
	return RegLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegLogic) Reg(req types.RegReq) (*types.CommonResp, error) {
	resp, err := l.svcCtx.UserRpc.Reg(
		l.ctx, &userclient.RegReq{
			Captcha: req.Captcha,
			Profile: &userclient.Profile{
				Kind:       req.Profile.Kind,
				Role:       req.Profile.Role,
				Phone:      req.Profile.Phone,
				Name:       req.Profile.Name,
				Nickname:   req.Profile.Nickname,
				Gender:     req.Profile.Gender,
				OpenId:     req.Profile.OpenId,
				From:       req.Profile.From,
				Password:   req.Profile.Password,
				IdNumber:   req.Profile.IdNumber,
				Organize:   req.Profile.Organize,
				Department: req.Profile.Department,
				JobTitle:   req.Profile.JobTitle,
				Avatar:     req.Profile.Avatar,
				Address:    req.Profile.Address,
			},
		},
	)

	authLogic := NewAuthLogic(l.ctx, l.svcCtx)
	token, err := authLogic.Auth()

	if err != nil {
		return nil, err
	}

	return &types.CommonResp{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    token,
	}, nil
}
