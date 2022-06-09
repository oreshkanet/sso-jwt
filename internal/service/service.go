package service

import (
	"context"
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
