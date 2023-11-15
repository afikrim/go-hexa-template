package api

import (
	"context"

	"github.com/afikrim/go-hexa-template/core/module"
	engine_v1 "github.com/afikrim/go-hexa-template/handler/api/pb/engine/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	Handler interface {
		engine_v1.EngineServiceServer

		Register(ctx context.Context, address string) error
	}

	handler struct {
		engine_v1.UnimplementedEngineServiceServer

		s   *grpc.Server
		mux *runtime.ServeMux
		svc module.BaseModule
	}
)

func New(s *grpc.Server, mux *runtime.ServeMux, svc module.BaseModule) Handler {
	return &handler{
		s:   s,
		mux: mux,
		svc: svc,
	}
}

func (h *handler) Register(ctx context.Context, address string) error {
	engine_v1.RegisterEngineServiceServer(h.s, h)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	return engine_v1.RegisterEngineServiceHandlerFromEndpoint(ctx, h.mux, address, opts)
}
