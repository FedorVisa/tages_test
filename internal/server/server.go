package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"tages_test/internal/errors"

	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"tages_test/gRPC/protobuf"
	"tages_test/internal/config"
)

type GRPCServer struct {
	protobuf.FileServiceServer
	config.Config
}

func Init() {

	unaryLimiter := rate.NewLimiter(100, 100)
	streamLimiter := rate.NewLimiter(10, 10)
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You are listening port: %s\n", cfg.Port)
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	setupServ := grpc.NewServer(
		grpc.StreamInterceptor(func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if streamLimiter.Allow() {
				return handler(srv, ss)
			}

			return errors.LimitSurpassionError{Method: info.FullMethod}.Error()

		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

			if unaryLimiter.Allow() {
				return handler(ctx, req)
			}

			return nil, errors.LimitSurpassionError{Method: info.FullMethod}.Error()

		}))
	grpcServ := &GRPCServer{}
	grpcServ.Filepath = cfg.Filepath
	protobuf.RegisterFileServiceServer(setupServ, grpcServ)
	if err = setupServ.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
