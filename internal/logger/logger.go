package logger

import (
	"log"

	"go.uber.org/fx"

	"github.com/dragonator/fx-golang/internal/service"
)

// Module -
var Module = fx.Options(
	fx.Provide(
		fx.Annotate(NewLogger, fx.As(new(service.ILogger))),
	),
)

// Logger -
type Logger struct{}

// NewLogger -
func NewLogger() *Logger {
	return &Logger{}
}

// Log -
func (l *Logger) Log(message string) {
	log.Println(message)
}
