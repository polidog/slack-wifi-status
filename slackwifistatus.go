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

			if config.LogOutput {
				log.Printf("change wifi: %s", status.WifiName)
			}

			wifi, err := config.FindWifi(status.WifiName)
			if err == nil {
				err = send(sender, status, wifi)
				if err != nil {
					if config.LogOutput {
						log.Println(err)
					}
					status.WifiName = ""
				}
			}
		}

		time.Sleep(time.Duration(config.Time) * time.Second)
	}
}

func send(sender sender.Slack, status status.Status, wifi config.Wifi) error {
	status.Message = wifi.Message
	status.Emoji = wifi.Emoji
	return sender.Send(status)
}
