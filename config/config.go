package config

import (
	"github.com/BurntSushi/toml"
	"github.com/go-errors/errors"
)

type Config struct {
	Slack    Slack    `toml:"slack"`
	WifiList WifiList `toml:"wifi"`
	isLoaded bool
}

func (c *Config) load(filename string) error {
	if c.isLoaded == false {
		_, err := toml.DecodeFile(filename, &c)
		if err != nil {
			return err
		}
		c.isLoaded = true
	}
	return nil
}

func (c *Config) FindWifi(wifiname string) (Wifi, error) {
	for _, wifi := range c.WifiList.Wifis {
		if wifi.Name == wifiname {
			return wifi, nil
		}
	}
	return Wifi{}, errors.New("wifiname not found:" + wifiname)
}

func NewConfig(filename string) (Config, error) {
	config := Config{
		isLoaded: false,
	}
	err := config.load(filename)
	return config, err
}
