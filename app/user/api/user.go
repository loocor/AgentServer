package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/rest/httpx"
	"google.golang.org/grpc/status"

	"agent/app/user/api/internal/config"
	"agent/app/user/api/internal/handler"
	"agent/app/user/api/internal/svc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(
		func(err error) (int, interface{}) {
			if st, ok := status.FromError(err); ok {
				return http.StatusOK, map[string]interface{}{
					"code":    st.Code(),
					"message": st.Message(),
				}
			} else {
				return http.StatusInternalServerError, nil
			}
		},
	)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
