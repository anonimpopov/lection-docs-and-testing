package ports

import (
	"context"

	"ru.mts.teta.tests_and_docs/internal/domain/models"
)

//go:generate mockery --name Auth

type Auth interface {
	Info(ctx context.Context, login string) (*models.User, error)
	Validate(ctx context.Context, tokens models.TokenPair) (string, error)
	Login(ctx context.Context, user, password string) (models.TokenPair, error)
}
