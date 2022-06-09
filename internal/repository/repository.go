package repository

import (
	"context"
	"github.com/oreshkanet/sso-jwt/internal/domain"
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
