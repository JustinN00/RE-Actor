package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type DiscordWebHook struct {
	URL string
}

func (dwh *DiscordWebHook) PostMessage(message string) (*http.Response, error) {
	// Discord only allows a max of 2000 characters in the content field.
	if len(message) > 2000 {
		return nil, fmt.Errorf("Message of %d characters is larger than the max of 2000 allowed for Discord", len(message))
	}
	body := fmt.Sprintf("{\"content\":\"%s\"}", message)
	client := http.Client{}
	response, err := client.Post(dwh.URL, "Application/json", strings.NewReader(body))
	return response, err
}
