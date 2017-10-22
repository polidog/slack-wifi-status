package sender

import "github.com/polidog/slack-wifi-status/status"

type Sender interface {
	Send(status status.Status)
}
