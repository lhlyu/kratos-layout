package service

import (
	"context"
	"errors"
	v1 "kratos-layout/gen/go/demo/v1"
	"time"

	"kratos-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServiceServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.SayHelloRequest) (*v1.SayHelloResponse, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	if in.Name == "jack" {
		return nil, errors.New("错误信息")
	}
	return &v1.SayHelloResponse{Message: "Hello " + g.Hello, Ts: time.Now().UnixMilli()}, nil
}
