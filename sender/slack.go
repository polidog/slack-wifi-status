package sender

import (
	"bytes"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/polidog/slack-wifi-status/config"
	"github.com/polidog/slack-wifi-status/status"
	"net/http"
	"net/url"
)

type Slack struct {
	token string
}

var apiUrl = "https://slack.com/api/users.profile.set"

func (s Slack) Send(status status.Status) error {
	if len(s.token) == 0 {
		return errors.New("token not found")
	}

	data := url.Values{}
	data.Set("token", s.token)
	data.Add("profile", status.ToJSON())

	client := &http.Client{}
	r, _ := http.NewRequest("POST", fmt.Sprintf("%s", apiUrl), bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := client.Do(r)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func NewSlackSender(slack config.Slack) Slack {
	return Slack{
		token: slack.Token,
	}
}
