package test_programs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func SusanTest() {

	fmt.Println("Starting Susan Testing function")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		reader, err := cli.ContainerLogs(ctx, ctr.ID, container.LogsOptions{
			ShowStdout: true,
			ShowStderr: true,
		})
		if err != nil {
			panic(err)
		}
		for {
			message, err := io.ReadAll(reader)
			if err != nil {
				panic(err)
			} else if len(message) == 0 {
				break
			}
			slog.Info(string(message))
		}
	}
}
