package utils

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/SusanHex/RE-Actor/actions"
	"github.com/SusanHex/RE-Actor/config"
)

func SelectAction(action_name string, app_config *config.Config) (actions.Action, error) {

	switch action_name {
	case "discord_webhook":
		return actions.DiscordWebHook{URL: app_config.DiscordWebHookURL}, nil
	default:
		return nil, fmt.Errorf(`"%s" does not match an action`, action_name)
	}
}

func PerformActionIfMatch(app_config *config.Config, action actions.Action, message []byte) error {
	match_indexes := app_config.CompiledPattern.FindSubmatchIndex(message)
	if len(match_indexes) == 0 {
		return nil
	}
	result := []byte{}
	result = app_config.CompiledPattern.Expand(result, []byte(app_config.Template), message, match_indexes)
	slog.Info(fmt.Sprintf(`Acting on "%v"`, result))
	err := action.Act(string(result))
	return err
}

func SetupLogger(app_config *config.Config) error {
	var log_level slog.Level
	add_source := false
	switch strings.ToUpper(app_config.LogLevel) {
	case "DEBUG":
		log_level = slog.LevelDebug
		add_source = true
	case "INFO":
		log_level = slog.LevelInfo
	case "WARNING":
		log_level = slog.LevelWarn
	case "ERROR":
		log_level = slog.LevelError
	default:
		return fmt.Errorf(`log level of "%s" does not match one of the following: DEBUG, INFO, WARNING, or ERROR`, app_config.LogLevel)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: add_source, Level: log_level}))
	slog.SetDefault(logger)
	return nil
}
