package main

import (
	"context"
	"github.com/oreshkanet/sso-jwt/internal/app"
	"golang.org/x/sync/errgroup"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	// FIXME: db, http, tokenizer

	// FIXME NewApp
	a := app.NewApp()

	ctxSignal, stop := signal.NotifyContext(ctx, syscall.SIGINT)

	g, ctxGroup := errgroup.WithContext(ctxSignal)
	g.Go(func() error {
		<-ctxGroup.Done()

		// Сигнал завершения приложения

		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer func() {
			_ = a.Stop(ctxTimeout)
			stop()
			cancel()
		}()

		<-ctxTimeout.Done()

		// Сервис остановлен

		return nil
	})

	// Старт сервиса

	g.Go(func() error {
		return a.Run(ctxGroup)
	})

	if err := g.Wait(); err != nil {
		//TODO: Errorf("errgroup.Wait")
	}
}
