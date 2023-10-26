package application

import (
	"context"
	"os"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"ru.mts.teta.tests_and_docs/internal/adapters/http"
	"ru.mts.teta.tests_and_docs/internal/adapters/postgres"
	"ru.mts.teta.tests_and_docs/internal/domain/auth"
)

var (
	s *http.Server
	logger *zap.Logger
)

func Start(ctx context.Context) {
	logger, _ = zap.NewProduction()

	pgconn := os.Getenv("PG_URL")

	db, err := postgres.New(ctx, pgconn)
	if err != nil {
		logger.Sugar().Fatalf("db init failed:", err)
	}
	authS := auth.New(db)
	
	s, err = http.New(logger.Sugar(), authS)
	if err != nil {
		logger.Sugar().Fatalf("http server creating failed:", err)
	}

	var g errgroup.Group
	g.Go(func() error {
		return s.Start()
	})

	logger.Sugar().Info("app is started")
	err = g.Wait()
	if err != nil {
		logger.Sugar().Fatalw("http server start failed", zap.Error(err))
	}
}

func Stop() {
	_ = s.Stop(context.Background())
	logger.Sugar().Info("app has stopped")
}
