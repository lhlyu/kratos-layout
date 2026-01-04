package main

import (
	"kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/lhlyu/kratos-easy/bootstrap"

	_ "go.uber.org/automaxprocs"
)

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return bootstrap.NewApp(
		logger,
		gs,
		hs,
	)
}

func main() {
	var bc conf.Conf

	bootstrap.Run(&bc, wireApp)
}
