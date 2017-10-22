package status

import (
	"encoding/json"
	"fmt"
	"github.com/polidog/slack-wifi-status/wifi"
)

type Status struct {
	WifiName string
	Message  string `json:"status_text"`
	Emoji    string `json:"status_emoji"`
}

func (s Status) ToJSON() string {
	json, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	fmt.Println(string(json))
	return string(json)
}

func (s *Status) Check() bool {
	name := wifi.GetWifiName()
	if s.WifiName != name {
		s.WifiName = name
		return true
	}
	return false
}
