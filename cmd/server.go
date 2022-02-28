package cmd

import (
	"context"
	"fmt"

	"wager-be/internal/app/server"
	"wager-be/internal/database"
	"wager-be/internal/repository"
	"wager-be/internal/service"
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

	db, err := database.NewPostgresConn(cfg.Database)
	if err != nil {
		logger.Error(err, "failed to init database conn")
		return
	}

	wagerRepo, err := repository.NewWagerRepository(db)
	if err != nil {
		logger.Error(err, "failed to init wager repo")
		return
	}

	purchaseRepo, err := repository.NewPurchaseRepository(db)
	if err != nil {
		logger.Error(err, "failed to init purchase repo")
	}

	wagerService := service.NewWagerService(wagerRepo, purchaseRepo)

	s := server.New(logger, cfg.Server, wagerService)

	s.Run(ctx)
}
