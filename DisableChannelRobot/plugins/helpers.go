package plugins

import (
	"errors"
	"fmt"

	"github.com/ALiwoto/DisableChannelRobot/DisableChannelRobot/core/logging"
	"github.com/ALiwoto/DisableChannelRobot/DisableChannelRobot/core/wotoConfig"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartTelegramBot() error {
	token := wotoConfig.GetBotToken()
	if len(token) == 0 {
		return errors.New("bot token is empty")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		RequestOpts: &gotgbot.RequestOpts{
			Timeout: 6 * gotgbot.DefaultTimeout,
		},
	})
	if err != nil {
		return err
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		MaxRoutines: MaxGoRoutines,
	})

	updater := ext.NewUpdater(dispatcher, &ext.UpdaterOpts{})
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: false,
	})
	if err != nil {
		return err
	}

	logging.Info(fmt.Sprintf("%s has started | ID: %d", b.Username, b.Id))

	LoadAllHandlers(dispatcher, wotoConfig.GetCmdPrefixes())

	updater.Idle()
	return nil
}
