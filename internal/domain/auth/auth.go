package auth

import (
	"context"
	"fmt"

	"ru.mts.teta.tests_and_docs/internal/domain/models"
	"ru.mts.teta.tests_and_docs/internal/ports"
)

type Service struct {
	db ports.UserStorage
}

func New(db ports.UserStorage) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Info(ctx context.Context, login string) (*models.User, error) {
	user, err :=  s.db.Get(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("get user info for login %s failed: %w", login, err)
	}
	return user, nil
}

func (s *Service) Validate(ctx context.Context, tokens models.TokenPair) (string, error) {
	// that is stub
	return "", nil
}

func (s *Service) Login(ctx context.Context, user, password string) (models.TokenPair, error) {
	// that is stub
	var tokens models.TokenPair
	return tokens, nil
}
