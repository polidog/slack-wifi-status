package wifi

import (
	"github.com/yelinaung/wifi-name"
	"log"
)

func GetWifiName() string {

	defer func() {
		err := recover()
		if err != nil {
			log.Println("wifi error")
			log.Println(err)
		}
	}()
	return wifiname.WifiName()
}
