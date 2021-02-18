package logic

import (
	"context"
	"strings"

	"google.golang.org/grpc/status"

	"agent/app/user/model"
	"agent/app/user/service/internal/svc"
	"agent/app/user/service/user"
	"agent/libs/codes"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegLogic {
	return &RegLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegLogic) Reg(in *user.RegReq) (*user.RegResp, error) {
	// Check parameter
	if len(strings.TrimSpace(in.Profile.Phone)) == 0 || len(strings.TrimSpace(in.Profile.IdNumber)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "reg touched, parameter error")
	}

	// Check exist by phone
	_, err := l.svcCtx.UserModel.FindOneByPhone(in.Profile.Phone)
	if err == nil {
		return nil, status.Errorf(codes.PhoneAlreadyExists, "reg touched, phone number exist")
	}

	// Check exist by Id number
	_, err = l.svcCtx.UserModel.FindOneByIdNumber(in.Profile.IdNumber)
	if err == nil {
		return nil, status.Errorf(codes.IDNumberAlreadyExists, "reg touched, id number exist")
	}

	// Insert data to DB
	_, err = l.svcCtx.UserModel.Insert(
		model.User{
			Kind:       in.Profile.Kind,
			State:      in.Profile.State,
			Role:       in.Profile.Role,
			Phone:      in.Profile.Phone,
			Name:       in.Profile.Name,
			Nickname:   in.Profile.Nickname,
			Gender:     in.Profile.Gender,
			OpenId:     in.Profile.OpenId,
			From:       in.Profile.From,
			Password:   in.Profile.Password,
			IdNumber:   in.Profile.IdNumber,
			Organize:   in.Profile.Organize,
			Department: in.Profile.Department,
			JobTitle:   in.Profile.JobTitle,
			Avatar:     in.Profile.Avatar,
			Address:    in.Profile.Address,
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.DBInsertError, "reg touched, db insert failed")
	}

	// Return result
	return &user.RegResp{
		Code:    0,
		Message: "reg succeed",
	}, nil
}
