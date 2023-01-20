package main

import (
	"context"
	"fmt"
	"github.com/oreshkanet/sso-jwt/internal/app"
	"github.com/oreshkanet/sso-jwt/internal/config"
	"github.com/oreshkanet/sso-jwt/pkg/database"
	"golang.org/x/sync/errgroup"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	// FIXME: db, http, tokenizer
	conf := config.NewConfig()

	// Создаём подключение к БД
	dbURL := fmt.Sprintf(
		"sqlserver://%s:%s@%s?database=%s",
		conf.MsSqlUser, conf.MsSqlPwd,
		conf.MsSqlServer, conf.MsSqlDb,
	)
	db, err := database.NewDBMsSQL(ctx, dbURL)
	if err != nil {
		log.Fatalf("%s", err.Error())
		return
	}
	defer db.Close()

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
