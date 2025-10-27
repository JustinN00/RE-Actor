package utils

import (
	"fmt"

	"github.com/SusanHex/RE-Actor/actions"
)

func SelectAction(action_name string) (actions.Action, error) {

	switch action_name {
	case "discord_webhook":
		return actions.DiscordWebHook{}, nil
	default:
		return nil, fmt.Errorf(`"%s" does not match an action`, action_name)
	}
}
