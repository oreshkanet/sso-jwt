package service

import (
	"context"
	"github.com/oreshkanet/sso-jwt/internal/repository"
)

type SoftwareAuthService interface {
	Auth(ctx context.Context, id string, pwd string) (string, error)
}

type ExternalClientAuthService interface {
	Auth(ctx context.Context, id string, token string, grantType string, scope string) (string, error)
}

type UserAuthService interface {
	Auth(ctx context.Context, id string, pwd string, scope string) (string, error)
}

type Service struct {
	Software       SoftwareAuthService
	ExternalClient ExternalClientAuthService
	User           UserAuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
