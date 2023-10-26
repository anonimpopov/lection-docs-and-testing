//go:build containers
// +build containers

package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"ru.mts.teta.tests_and_docs/internal/adapters/http"
	"ru.mts.teta.tests_and_docs/internal/adapters/postgres"
	"ru.mts.teta.tests_and_docs/internal/domain/auth"
)

func TestTestcontainers(t *testing.T) {
	suite.Run(t, new(TestcontainersSuite))
}

type TestcontainersSuite struct {
	suite.Suite

	server      *http.Server
	pgContainer testcontainers.Container
}

const (
	dbName = "messages"
	dbUser = "admin"
	dbPass = "admin"
)

func (suite *TestcontainersSuite) SetupSuite() {
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()

	dbContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:11",
			ExposedPorts: []string{"5432"},
			Env: map[string]string{
				"POSTGRES_DB":       dbName,
				"POSTGRES_USER":     dbUser,
				"POSTGRES_PASSWORD": dbPass,
			},
			WaitingFor: wait.ForLog("database system is ready to accept connections"),
			SkipReaper: true,
			AutoRemove: true,
		},
		Started: true,
	})
	suite.Require().NoError(err)

	// with a second delay migrations work properly
	time.Sleep(time.Second * 5)

	ip, err := dbContainer.Host(ctx)
	suite.Require().NoError(err)
	port, err := dbContainer.MappedPort(ctx, "5432")
	suite.Require().NoError(err)

	cfg := &pgx.ConnConfig{
		Config: pgconn.Config{
			Host:     ip,
			Port:     uint16(port.Int()),
			Database: dbName,
			User:     dbUser,
			Password: dbPass,
		},
	}

	connString := fmt.Sprintf(`postgres://%s:%s@%s:%d/%s?sslmode=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		"disable",
	)

	config, err := pgxpool.ParseConfig(connString)
	suite.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	suite.Require().NoError(err)

	auth := auth.New(&postgres.Database{
		DB: pool,
	})

	server, err := http.New(logger.Sugar(), auth)
	suite.Require().NoError(err)
	go func() {
		server.Start()
	}()
	suite.server = server
	suite.pgContainer = dbContainer

	suite.T().Log("Suite setup is done")
}

func (s *TestcontainersSuite) TearDownSuite() {
	s.pgContainer.Terminate(context.Background())
	s.T().Log("Suite stop is done")
}
