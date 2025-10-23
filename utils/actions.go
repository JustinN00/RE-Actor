package utils

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

type DiscordWebHook struct {
	Id    string
	Token string
}

func (dwh *DiscordWebHook) post_message(message string) {
	url := fmt.Sprintf("https://discord.com/api/webhook/%s/%s", dwh.id, dwh.token)
	body := fmt.Sprintf("{\"content\":\"%s\"}", message)
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	body_content := []byte{}
	_, err = req.Response.Body.Read(body_content)
	if err != nil {
		panic(err)
	}
	slog.Debug("%s")
}
