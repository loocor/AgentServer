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

type ProfileUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProfileUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileUpdateLogic {
	return &ProfileUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProfileUpdateLogic) ProfileUpdate(in *user.ProfileUpdateReq) (*user.ProfileUpdateResp, error) {
	// Check parameter
	if in.Profile.Id == 0 || len(strings.TrimSpace(in.Profile.Phone)) == 0 || len(strings.TrimSpace(in.Profile.IdNumber)) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "profile update touched, parameter error")
	}

	// check exist by id
	if _, err := l.svcCtx.UserModel.FindOne(in.Profile.Id); err != nil {
		return nil, status.Errorf(codes.UserNotFound, "profile update touched, user not exist")
	}

	// check new phone number if exist
	if _user, err := l.svcCtx.UserModel.FindOneByPhone(in.Profile.Phone); err == nil && _user.Id != in.Profile.Id {
		return nil, status.Errorf(codes.PhoneAlreadyExists, "profile update touched, new phone number exist")
	}

	// check new id number if exist
	if _user, err := l.svcCtx.UserModel.FindOneByIdNumber(in.Profile.IdNumber); err == nil && _user.Id != in.Profile.Id {
		return nil, status.Errorf(codes.IDNumberAlreadyExists, "profile update touched, new id number exist")
	}

	// Update data to DB
	err := l.svcCtx.UserModel.Update(
		model.User{
			Id:         in.Profile.Id,
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
		return nil, status.Errorf(codes.DBUpdateError, "reg touched, db update failed")
	}

	return &user.ProfileUpdateResp{
		Code:    0,
		Message: "profile update succeed",
	}, nil
}
