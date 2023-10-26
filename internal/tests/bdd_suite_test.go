//go:build bdd
// +build bdd

package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/zap"
	"ru.mts.teta.tests_and_docs/internal/adapters/http"
	"ru.mts.teta.tests_and_docs/internal/adapters/postgres"
	"ru.mts.teta.tests_and_docs/internal/domain/auth"
)

func TestBDD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BDD Suite")
}

const (
	defaultTimeoutSec = 10
)



var (
	dbContainer testcontainers.Container
	server      *http.Server
)

var _ = BeforeSuite(func() {
	const (
		dbName = "messages"
		dbUser = "admin"
		dbPass = "admin"
	)

	var err error
	logger, _ := zap.NewDevelopment()
	ctx := context.Background()

	dbContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
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
	Expect(err).To(BeNil())

	// with a second delay migrations work properly
	time.Sleep(time.Second * 5)

	ip, err := dbContainer.Host(ctx)
	Expect(err).To(BeNil())
	port, err := dbContainer.MappedPort(ctx, "5432")
	Expect(err).To(BeNil())

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
	Expect(err).To(BeNil())
	pool, err := pgxpool.ConnectConfig(ctx, config)
	Expect(err).To(BeNil())

	auth := auth.New(&postgres.Database{
		DB: pool,
	})

	server, err = http.New(logger.Sugar(), auth)
	Expect(err).To(BeNil())
	go func() {
		server.Start()
	}()
}, defaultTimeoutSec)

var _ = AfterSuite(func() {
	if dbContainer != nil {
		err := dbContainer.Terminate(context.Background())
		Expect(err).NotTo(HaveOccurred())
	}
}, defaultTimeoutSec)
