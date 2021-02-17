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

func (l *LoginLogic) Login(in *user.LoginReq) (*user.CommonResp, error) {
	// Check parameter
	if len(strings.TrimSpace(in.Phone)) == 0 || len(strings.TrimSpace(in.Password)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "login touched, parameter error")
	}

	// Query data
	switch info, err := l.svcCtx.UserModel.FindOneByPhone(in.Phone); err {
	case nil:
		if info.Password != in.Password {
			return nil, status.Errorf(codes.PasswordWrong, "login touched, password error")
		}
	case model.ErrNotFound:
		return nil, status.Errorf(codes.UserNotFound, "login touched, user not exist")
	default:
		return nil, status.Errorf(codes.Unknown, "login touched, unknown error")
	}

	// Return result
	return &user.CommonResp{
		Code:    0,
		Message: "login Succeed",
		Data:    nil, // TODO: Return profile with id
	}, nil
}
