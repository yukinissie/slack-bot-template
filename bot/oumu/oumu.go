package oumu

import (
	"github.com/slack-go/slack"
)
func Gaeshi(api *slack.Client, channel string, text string, botId string) {
	if botId != "" {
		return
	}
	api.PostMessage(channel, slack.MsgOptionText(text, true))
}
