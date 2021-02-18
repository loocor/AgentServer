package logic

import (
	"context"
	"strings"

	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/status"

	"agent/app/user/model"
	"agent/app/user/service/internal/svc"
	"agent/app/user/service/user"
	"agent/libs/codes"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// Check parameter
	if len(strings.TrimSpace(in.Phone)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "login failed, parameter error")
	}

	// Query data
	record, err := l.svcCtx.UserModel.FindOneByPhone(in.Phone)

	switch err {
	case nil:
		if record.Password != in.Password {
			return nil, status.Errorf(codes.PasswordWrong, "login failed, password error")
		}
	case model.ErrNotFound:
		return nil, status.Errorf(codes.UserNotFound, "login failed, user not exist")
	default:
		return nil, status.Errorf(codes.Unknown, "login failed, unknown error")
	}

	// Check state
	switch record.State {
	case 0:
	case 1:
	case 2:
		return nil, status.Errorf(codes.UserFreeze, "login failed, account freeze")
	case 3:
		return nil, status.Errorf(codes.UserRemoved, "login failed, account removed")
	}

	// Return result
	return &user.LoginResp{
		Code:    0,
		Message: "login Succeed",
	}, nil
}
