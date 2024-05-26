package logger

import (
	"github.com/rs/zerolog"
	"os"
	"rat/internal/core/config"
)

var Logger zerolog.Logger

func Init() {
	if config.Get().Logger.ErrorLvl {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	Logger = zerolog.New(os.Stderr)
}

func Get() *zerolog.Logger {
	if &Logger == nil {
		Init()
	}
	return &Logger
}
