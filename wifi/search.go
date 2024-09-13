package wifi

import (
	"log"
	"os/exec"
	"regexp"
	"runtime"
)

func GetWifiName() string {

	defer func() {
		err := recover()
		if err != nil {
			log.Println("wifi error")
			log.Println(err)
		}
	}()
	return ssid()
}

// Support only drawin（macOS）
func ssid() string {
	var ssid string
	switch runtime.GOOS {
	case "darwin":
		ssid = forDarwin()
	}
	return ssid
}

func forDarwin() string {
	device_output, err := exec.Command("networksetup", "-listallhardwareports").Output()
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(`Hardware Port: Wi-Fi\nDevice: (.+)\n`)
	wifi_device := string(r.FindSubmatch(device_output)[1])
	ssid_output, err := exec.Command("networksetup", "-getairportnetwork", wifi_device).Output()
	if err != nil {
		panic(err)
	}
	r = regexp.MustCompile(`Current Wi-Fi Network: (.+)`)
	if !r.Match(ssid_output) {
		panic("wifi is not connected")
	}
	ssid := string(r.FindSubmatch(ssid_output)[1])

	return ssid
}
