package antiChannelPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadHandlers(d *ext.Dispatcher, t []rune) {
	d.AddHandler(handlers.NewMessage(antiChannelFilter, checkAntiChannel))
}
