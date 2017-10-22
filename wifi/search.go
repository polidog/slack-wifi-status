package wifi

import "github.com/yelinaung/wifi-name"

func GetWifiName() string {
	return wifiname.WifiName()
}
