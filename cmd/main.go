package main

import (
	"context"
	"easymvp_api/internal/app"
	"easymvp_api/internal/database"
	"easymvp_api/internal/log"
	"easymvp_api/internal/swagger"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"time"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// set timezone to be utc
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc

	app.LoadEnv()

	server := fx.New(
		log.Module, // NOTE: comment out this line to use other zap logger
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger}
		}),
		database.Module,
		app.Module,
		swagger.Module,
		fx.Invoke(
			func(s *app.App, logger *zap.Logger) {
				addr := fmt.Sprintf("%s:%s", s.HttpConfig.Host, s.HttpConfig.Port)
				go func() {
					_ = s.Run(addr)
				}()
			},
		),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Start(startCtx); err != nil {
		panic(err)
	}
	<-server.Wait()
}
