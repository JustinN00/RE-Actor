package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
)

type DiscordWebHook struct {
	URL string
}

func (dwh DiscordWebHook) PostMessage(message string) (*http.Response, error) {
	// Discord only allows a max of 2000 characters in the content field.
	if len(message) > 2000 {
		return nil, fmt.Errorf("message of %d characters is larger than the max of 2000 allowed for Discord", len(message))
	}

	body_struct := struct {
		Content string `json:"content"`
	}{
		Content: message,
	}
	body_bytes, err := json.Marshal(body_struct)
	body := string(body_bytes)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	response, err := client.Post(dwh.URL, "Application/json", strings.NewReader(body))
	slog.Debug("Discord Webhook Response:", "Status Code", response.StatusCode, "Status", response.Status)
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		response_content, err := io.ReadAll(response.Body)
		if err != nil {
			slog.Error("Error reading response body: ", "Error", err)
		}
		slog.Debug(fmt.Sprintf(`Response Content: "%s"`, response_content))
	}
	return response, err
}

func (dwh DiscordWebHook) Act(message string) error {
	_, err := dwh.PostMessage(message)
	return err
}
