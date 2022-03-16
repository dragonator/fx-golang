package main

import (
	"go.uber.org/fx"

	"github.com/dragonator/fx-golang/internal/logger"
	"github.com/dragonator/fx-golang/internal/service"
)

func main() {
	var a *logger.Logger
	// a = logger.NewLogger()
	// a.WithPrefix("testing-replace: ")

	app := fx.New(
		logger.Module,
		fx.Invoke(service.RunFirst),
		fx.Invoke(service.RunSecond),
		fx.Invoke(service.RunThird),
		fx.Invoke(service.RegisterFirstHook),
		fx.Invoke(service.RegisterSecondHook),
		fx.Decorate(func(logger service.ILogger) service.ILogger {
			logger.WithPrefix("testing-decorate: ")
			return logger
		}),
		// fx.Replace(
		// 	fx.Annotate(a, fx.As(new(service.ILogger))),
		// ),
		fx.Populate(&a),
		fx.Invoke(service.RunWithPAramStruct),
	)
	app.Run()
}
