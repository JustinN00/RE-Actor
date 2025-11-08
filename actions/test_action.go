package actions

import "log/slog"

type TestAction struct {}

func (ta TestAction) Act(message string) error {
	slog.Info("Test Action:", "message", message)
	return nil
}
