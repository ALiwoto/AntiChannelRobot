package main

import (
	"log"

	"github.com/ALiwoto/DisableChannelRobot/src/core/logging"
	"github.com/ALiwoto/DisableChannelRobot/src/core/wotoConfig"
	"github.com/ALiwoto/DisableChannelRobot/src/plugins"
)

func main() {
	_, err := wotoConfig.LoadConfig()
	if err != nil {
		log.Fatal("Error parsing config file", err)
	}

	f := logging.LoadLogger()
	if f != nil {
		defer f()
	}

	err = plugins.StartTelegramBot()
	if err != nil {
		logging.Fatal("Failed to start the bot bot: ", err)
	}
}
