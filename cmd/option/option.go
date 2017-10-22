package option

import (
	"os"
	"runtime"
)

type Option struct {
	Version  bool   `long:"version" description:"Show version"`
	Filename string `short:"c" long:"config" description:"Filename file name."`
}

func (o Option) GetFilename() string {
	if len(o.Filename) > 0 {
		return o.Filename
	}
	return homeDir() + "/.slack-wifi-status.toml"
}

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
