package app

import (
	"github.com/polidog/slack-wifi-status/config"
	"github.com/polidog/slack-wifi-status/sender"
	"github.com/polidog/slack-wifi-status/status"
	"log"
	"time"
)

func Run(config config.Config) {
	sender := sender.NewSlackSender(config.Slack)
	status := status.Status{}

	for {

		if status.Check() {
			wifi, err := config.FindWifi(status.WifiName)
			if err == nil {
				send(sender, status, wifi)
			}
		}

		time.Sleep(time.Duration(60) * time.Second)
	}
}

func send(sender sender.Slack, status status.Status, wifi config.Wifi) {
	status.Message = wifi.Message
	status.Emoji = wifi.Emoji
	err := sender.Send(status)
	if err != nil {
		log.Println(err)
	}
}
