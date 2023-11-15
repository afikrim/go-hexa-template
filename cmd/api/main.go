package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/afikrim/go-hexa-template/config"
	handler "github.com/afikrim/go-hexa-template/handler/api"
	"github.com/afikrim/go-hexa-template/repo"
	"github.com/afikrim/go-hexa-template/service"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg := config.New(
		config.WithFile(".env"),
		config.WithFile(".env.local"),
		config.WithEtcd("afikrim/go-hexa-template", []string{"192.168.100.12:2379"}),
		config.WithEnv(),
	)

	db, err := cfg.Database.Init()
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	repo := repo.New(db)
	if cfg.Database.Migration {
		if err := repo.Migrate(); err != nil {
			panic(fmt.Errorf("failed to migrate database: %w", err))
		}
	}

	svc := service.New(repo)

	mux := runtime.NewServeMux()
	grpcServer := grpc.NewServer()
	httpServer := http.Server{
		Handler: mux,
	}

	h := handler.New(grpcServer, mux, svc)

	grpcListen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Application.GRPCPort))
	if err != nil {
		panic(fmt.Errorf("failed to listen grpc: %w", err))
	}

	httpListen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Application.HTTPPort))
	if err != nil {
		panic(fmt.Errorf("failed to listen http: %w", err))
	}

	go func() {
		if err := grpcServer.Serve(grpcListen); err != nil {
			panic(fmt.Errorf("failed to serve grpc: %w", err))
		}
	}()

	go func() {
		if err := httpServer.Serve(httpListen); err != nil {
			panic(fmt.Errorf("failed to serve http: %w", err))
		}
	}()

	if err := h.Register(ctx, fmt.Sprintf("localhost:%d", cfg.Application.GRPCPort)); err != nil {
		panic(fmt.Errorf("failed to register handler: %w", err))
	}

	killCh := make(chan os.Signal, 1)
	signal.Notify(killCh, syscall.SIGTERM, syscall.SIGINT)
	<-killCh

	grpcServer.GracefulStop()
	httpServer.Shutdown(ctx)
}
