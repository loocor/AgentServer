package logic

import (
	"context"
	"strings"

	"github.com/jinzhu/copier"
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

func (l *ProfileViewLogic) ProfileView(in *user.ProfileViewReq) (*user.ProfileViewResp, error) {
	// Check parameter
	if len(strings.TrimSpace(in.Phone)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "profile view touched, parameter error")
	}

	// Try to get record
	resp, err := l.svcCtx.UserModel.FindOneByPhone(in.Phone)
	if err != nil {
		return nil, status.Errorf(codes.UserNotFound, "profile view touched, user not exist")
	}

	// Copy record to profile
	profile := &user.Profile{}
	if err := copier.Copy(&profile, &resp); err != nil {
		logx.Error(err)
	}

	// Correct timestamp
	profile.CreateTime = resp.CreateTime.Unix()
	profile.UpdateTime = resp.UpdateTime.Unix()

	return &user.ProfileViewResp{
		Code:    0,
		Message: "profile view succeed",
		Profile: profile,
	}, nil
}
