package ports

import (
	"context"

	"ru.mts.teta.tests_and_docs/internal/domain/models"
)

//go:generate mockery --name UserStorage

type UserStorage interface {
	Get(ctx context.Context, login string) (*models.User, error)
}
