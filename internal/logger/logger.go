package logger

import (
	"log"
	"os"

	"go.uber.org/fx"

	"github.com/dragonator/fx-golang/internal/service"
)

var a *Logger = NewLogger()

// Module -
var Module = fx.Options(
	// fx.Provide(
	// 	fx.Annotate(NewLogger, fx.As(new(service.ILogger))),
	// ),
	fx.Supply(
		a,
		// annotation overwrites original type. in order to provide the original type one must provide/supply it separately
		fx.Annotate(a, fx.As(new(service.ILogger))),
	),
)

// Logger -
type Logger struct {
	log *log.Logger
}

// NewLogger -
func NewLogger() *Logger {
	return &Logger{
		log: log.New(os.Stdout, "", 0),
	}
}

// Log -
func (l *Logger) Log(message string) {
	l.log.Println(message)
}

// WithPrefix -
func (l *Logger) WithPrefix(prefix string) {
	l.log.SetPrefix(prefix)
}
