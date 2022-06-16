package app

import (
	"context"
	"github.com/oreshkanet/sso-jwt/internal/delivery/api"
	"github.com/oreshkanet/sso-jwt/internal/repository"
	"github.com/oreshkanet/sso-jwt/internal/service"
	"github.com/oreshkanet/sso-jwt/pkg/database"
	"github.com/oreshkanet/sso-jwt/pkg/tokenizer"
	"golang.org/x/sync/errgroup"
	"net/http"
)

type App struct {
	db    database.DB
	http  *http.Server
	token *tokenizer.Tokenizer
}

func NewApp(
	db database.DB,
	http *http.Server,
	token *tokenizer.Tokenizer,
) *App {
	return &App{
		db:    db,
		http:  http,
		token: token,
	}
}

func (a App) Run(ctx context.Context) error {
	s := service.NewService(
		repository.NewRepository(a.db))

	httpApi := api.NewApi(&api.Config{
		Srv:          a.http,
		Token:        a.token,
		SoftwareAuth: s.Software,
		ExternalAuth: s.ExternalClient,
		UserAuth:     s.User,
	})

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpApi.Run()
	})

	//TODO: Worker

	return g.Wait()
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.http.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
