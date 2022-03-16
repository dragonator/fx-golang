package service

import (
	"context"

	"go.uber.org/fx"
)

// FirstOnStart -
func FirstOnStart(logger ILogger) {
	logger.Log("FIRST OnStart hook")
}

// SecondOnStart -
func SecondOnStart(logger ILogger) {
	logger.Log("SECOND OnStart hook")
}

// RegisterFirstHook -
func RegisterFirstHook(lifesycle fx.Lifecycle, logger ILogger) {
	lifesycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("OnStart hook 1-1")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("OnStop hook 1-1")
			return nil
		},
	})
	lifesycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("OnStart hook 1-2")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("OnStop hook 1-2")
			return nil
		},
	})
}

// RegisterSecondHook -
func RegisterSecondHook(lifesycle fx.Lifecycle, logger ILogger) {
	lifesycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("OnStop hook 1-1")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("OnStop hook 2-1")
			return nil
		},
	})
	lifesycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Log("OnStart hook 2-2")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Log("OnStop hook 2-2")
			return nil
		},
	})
}
