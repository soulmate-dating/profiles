package grpc

import (
	"context"
	"github.com/soulmate-dating/profiles/internal/ports/http"
	"log"
	"net"
	"os"

	grpcProm "github.com/grpc-ecosystem/go-grpc-prometheus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/soulmate-dating/profiles/internal/app"
	"github.com/soulmate-dating/profiles/internal/config"
	"github.com/soulmate-dating/profiles/internal/graceful"
)

const MB = 1024 * 1024

func Run(ctx context.Context, cfg config.Config, app app.App) {
	lis, err := net.Listen(cfg.API.Network, cfg.API.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	svc := NewService(app)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			UnaryLoggerInterceptor,
			UnaryRecoveryInterceptor(),
		),
		grpc.MaxRecvMsgSize(cfg.API.MaxReceiveSize*MB),
		grpc.MaxSendMsgSize(cfg.API.MaxSendSize*MB),
		grpc.StreamInterceptor(grpcProm.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpcProm.UnaryServerInterceptor),
	)
	RegisterProfileServiceServer(grpcServer, svc)

	grpcProm.Register(grpcServer)
	s := http.NewServer(cfg.Metrics.Address)

	eg, ctx := errgroup.WithContext(ctx)
	sigQuit := make(chan os.Signal, 1)
	eg.Go(graceful.CaptureSignal(ctx, sigQuit))
	eg.Go(RunGRPCServerGracefully(ctx, lis, grpcServer))
	eg.Go(http.RunServer(ctx, s))

	if err := eg.Wait(); err != nil {
		log.Printf("gracefully shutting down the servers: %s\n", err.Error())
	}
	log.Println("servers were successfully shutdown")
}
