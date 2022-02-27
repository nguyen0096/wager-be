package cmd

import (
	"context"
	"fmt"

	"wager-be/internal/app/server"
	"wager-be/pkg/config"
	"wager-be/pkg/log"
)

func runServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config. err: %w", err))
	}

	logger := log.NewLogger()

	s := server.New(logger, cfg.Server)

	s.Run(ctx)
}
