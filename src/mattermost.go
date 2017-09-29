package main

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
)

func MattermostNotify(channel string, message string) {
	payload := slack.Payload {
		Text: message,
		Username: "Release Manager",
		Channel: "#" + channel,
		IconUrl: "https://gitlab.com/uploads/-/system/project/avatar/430285/ship-it-squirrel.png",
	}
	err := slack.Send(ConfigFile().MattermostWebhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}