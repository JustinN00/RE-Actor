package utils

import (
	"fmt"
	"log/slog"

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
	slog.Info("Acting on", "message", string(result))
	err := action.Act(string(result))
	return err
}

