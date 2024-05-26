package main

import (
	"rat/internal/core/config"
	"rat/internal/core/logger"
)

func init() {
	config.Init()
	logger.Init()
}

func main() {
	logger.Get().Info().Timestamp().Str("application", config.Get().AppInfo.AppName).Msg("app started")
}
