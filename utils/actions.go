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
	body := fmt.Sprintf("{\"content\":\"%s\"}", message)
	client := http.Client{}
	response, err := client.Post(dwh.URL, "Application/json", strings.NewReader(body))
	return response, err
}
