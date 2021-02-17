package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"

	"agent/app/user/api/internal/config"
	"agent/app/user/api/internal/middleware"
	"agent/app/user/service/userclient"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   userclient.User
	UserCheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		UserCheck: middleware.NewUserCheckMiddleware().Handle,
	}
}
