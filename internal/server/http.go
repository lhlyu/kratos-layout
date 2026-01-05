package server

import (
	v1 "kratos-layout/api/demo/v1"
	"kratos-layout/internal/conf"
	"kratos-layout/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lhlyu/kratos-easy/httpx"
	"github.com/lhlyu/kratos-easy/middlewares/header"
	"github.com/lhlyu/kratos-easy/middlewares/logging"
	"github.com/lhlyu/kratos-easy/middlewares/validate"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Conf, logger log.Logger, greeter *service.GreeterService) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.ProtoValidate(),
			header.Header(),
		),
		http.ResponseEncoder(httpx.EncodeResponse),
		http.ErrorEncoder(httpx.EncodeError),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterServiceHTTPServer(srv, greeter)
	return srv
}
