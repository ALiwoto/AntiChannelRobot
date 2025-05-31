package plugins

import (
	"github.com/ALiwoto/DisableChannelRobot/src/plugins/antiChannelPlugin"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func LoadAllHandlers(d *ext.Dispatcher, triggers []rune) {
	antiChannelPlugin.LoadHandlers(d, triggers)
}
