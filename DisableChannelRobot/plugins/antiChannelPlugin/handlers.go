package antiChannelPlugin

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func checkAntiChannel(bot *gotgbot.Bot, ctx *ext.Context) error {
	message := ctx.EffectiveMessage
	senderChat := message.SenderChat
	chat := ctx.EffectiveChat

	_, _ = bot.DeleteMessage(chat.Id, message.MessageId)
	_, _ = bot.BanChatSenderChat(chat.Id, senderChat.Id, nil)

	// don't let another handlers to be executed
	return ext.EndGroups
}

func antiChannelFilter(msg *gotgbot.Message) bool {
	if msg.From != nil {
		if msg.From.Id == 1087968824 {
			return false
		}
	}
	return msg.SenderChat != nil &&
		!msg.IsAutomaticForward
}
