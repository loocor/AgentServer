package logic

import (
	"context"
	"strings"

	"google.golang.org/grpc/status"

	"agent/app/user/service/internal/svc"
	"agent/app/user/service/user"
	"agent/libs/codes"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProfileViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProfileViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileViewLogic {
	return &ProfileViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProfileViewLogic) ProfileView(in *user.ProfileViewReq) (*user.CommonResp, error) {
	// Check parameter
	if len(strings.TrimSpace(in.Phone)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "profile view touched, parameter error")
	}

	// Try to get record
	_, err := l.svcCtx.UserModel.FindOneByPhone(in.Phone)
	if err != nil {
		return nil, status.Errorf(codes.UserNotFound, "profile view touched, user not exist")
	}

	return &user.CommonResp{
		Code:    0,
		Message: "profile view touched",
		Data:    nil, // TODO: Return profile info
	}, nil
}
