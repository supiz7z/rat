package logger

import (
	"github.com/rs/zerolog"
	"os"
	"rat/internal/core/config"
)

var Logger zerolog.Logger

func Init() {
	Logger = zerolog.New(os.Stdout)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if config.Get().Logger.ErrorLvl {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

func Get() *zerolog.Logger {
	if &Logger == nil {
		Init()
	}
	return &Logger
}
