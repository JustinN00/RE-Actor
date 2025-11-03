/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"

	"github.com/SusanHex/RE-Actor/config"
	"github.com/SusanHex/RE-Actor/utils"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/spf13/viper"
)

func main() {
	viper_instance := viper.NewWithOptions()
	app_config, err := config.GetConfigFromViper(viper_instance)
	if err != nil {
		panic(err)
	}
	slog.Info("Current config:", "Container name", app_config.ContainerName, "Action", app_config.ActionName, "Pattern", app_config.Pattern, "Template", app_config.Template)
	action, err := utils.SelectAction(app_config.ActionName, app_config)
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf(`Found action "%T"`, action))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	container_name_filter := filters.NewArgs(filters.KeyValuePair{Key: "name", Value: app_config.ContainerName})

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{Filters: container_name_filter})
	if err != nil {
		panic(err)
	} else if len(containers) == 0 {
		panic(fmt.Sprintf(`Could not find a container named "%s"`, app_config.ContainerName))
	} else if len(containers) > 1 {
		panic(fmt.Sprintf(`Container name: "%s" matched %s containers. Please ensure that the container name is unique to one container.`, app_config.ContainerName, len(containers)))
	}
	ctr := containers[0]
	reader, err := cli.ContainerLogs(ctx, ctr.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		panic(err)
	}
	for {
		message, err := io.ReadAll(reader)
		if len(message) == 0 {
			continue
		}
		if err != nil {
			panic(err)
		}
		err = utils.PerformActionIfMatch(app_config, action, message[8:])
		if err != nil {
			panic(err)
		}
	}
}
