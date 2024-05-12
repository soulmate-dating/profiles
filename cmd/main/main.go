package main

import (
	"context"
	"log"

	"github.com/soulmate-dating/profiles/internal/app"
	"github.com/soulmate-dating/profiles/internal/config"
	"github.com/soulmate-dating/profiles/internal/ports/grpc"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	appSvc := app.New(ctx, cfg)
	grpc.Run(ctx, cfg, appSvc)
}
