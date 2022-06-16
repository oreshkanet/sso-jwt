package repository

import (
	"context"
	"github.com/oreshkanet/sso-jwt/internal/domain"
	"github.com/oreshkanet/sso-jwt/pkg/database"
)

type SoftwareRepository interface {
	FindById(ctx context.Context, id string) (*domain.Software, error)
}

type ExternalClientRepository interface {
	FindById(ctx context.Context, id string) (*domain.ExternalClient, error)
}

type UserRepository interface {
	FindById(ctx context.Context, id string) (*domain.User, error)
}

type Repository struct {
}

func NewRepository(db database.DB) *Repository {
	return &Repository{}
}
