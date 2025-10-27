/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log/slog"

	"github.com/SusanHex/RE-Actor/config"
	"github.com/SusanHex/RE-Actor/utils"
	"github.com/spf13/viper"
)

func main() {
	app_config := config.Config{}
	viper_instance := viper.NewWithOptions()
	viper_instance.BindEnv("pattern")
	viper_instance.BindEnv("template")
	viper_instance.BindEnv("action_name")
	viper_instance.AutomaticEnv()
	err := viper_instance.UnmarshalExact(&app_config)
	if err != nil {
		panic(err)
	}
	slog.Info("Current config: ", "Action", app_config.ActionName, "Pattern:", app_config.Pattern, "Template", app_config.Template)
	action, err := utils.SelectAction(app_config.ActionName)
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf(`Found action "%T"`, action))
}
