package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"

	"agent/app/user/api/internal/svc"
	"agent/app/user/api/internal/types"
	"agent/app/user/service/userclient"
)

type ProfileUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProfileUpdateLogic {
	return ProfileUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileUpdateLogic) ProfileUpdate(req types.ProfileUpdateReq) (*types.CommonResp, error) {
	resp, err := l.svcCtx.UserRpc.ProfileUpdate(
		l.ctx, &userclient.ProfileUpdateReq{
			Profile: &userclient.Profile{
				Id:         req.Profile.Id,
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

	if err != nil {
		return nil, err
	}

	return &types.CommonResp{
		Code:    resp.Code,
		Message: resp.Message,
		Data:    resp.Data,
	}, nil
}
