package main

import (
	"go.uber.org/fx"

	"github.com/dragonator/fx-golang/internal/logger"
	"github.com/dragonator/fx-golang/internal/service"
)

func main() {
	app := fx.New(
		logger.Module,
		fx.Invoke(service.RunFirst),
		fx.Invoke(service.RunSecond),
		fx.Invoke(service.RunThird),
		fx.Invoke(service.RegisterFirstHook),
		fx.Invoke(service.RegisterSecondHook),
	)
	app.Run()
}
