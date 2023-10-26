//go:build testing || testify
// +build testing testify

package tests

import (
	"context"
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"
	"ru.mts.teta.tests_and_docs/internal/adapters/http"
	"ru.mts.teta.tests_and_docs/internal/adapters/postgres"
	"ru.mts.teta.tests_and_docs/internal/domain/auth"
)

var (
	s *http.Server
)

const (
	pgURL = "postgres://admin:admin@localhost:5432/auth"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup tests")
	logger, _ := zap.NewDevelopment()

	db, err := postgres.New(context.TODO(), pgURL)
	s, err = http.New(logger.Sugar(), auth.New(db))
	if err != nil {
		logger.Fatal("http server creation failed", zap.Error(err))
	}
	go func() {
		err := s.Start()
		logger.Fatal("http server start failed", zap.Error(err))
	}()
	fmt.Println("Setup completed")
	os.Exit(m.Run())
}
