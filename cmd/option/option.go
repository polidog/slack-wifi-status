package option

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Option struct {
	Version  bool   `long:"version" description:"Show version"`
	Filename string `short:"c" long:"config" description:"Filename file name."`
}

func (o Option) GetFilename() string {
	homedir := homeDir()
	if len(o.Filename) > 0 {
		path, err := filepath.Abs(o.Filename)
		if err != nil {
			log.Fatal(err)
		}

		wd, _ := os.Getwd()
		return strings.Replace(path, fmt.Sprintf("%s/~", wd), homedir, 1)
	}

	return filepath.Join(homedir, ".slack-wifi-status.toml")
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

func replaceUserDirectry(path string) string {
	return path
}
