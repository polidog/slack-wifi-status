package config

type WifiList struct {
	Wifis []Wifi `toml:"list"`
}

type Wifi struct {
	Name    string `toml:"name"`
	Message string `toml:"message"`
	Emoji   string `toml:"emoji"`
}
