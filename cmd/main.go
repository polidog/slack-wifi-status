package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"github.com/polidog/slack-wifi-status"
	"github.com/polidog/slack-wifi-status/cmd/option"
	"github.com/polidog/slack-wifi-status/config"
	"log"
)

const version = "0.0.1"

var cmdOption option.Option

func main() {
	parser := flags.NewParser(&cmdOption, flags.Default)
	parser.Name = "slack-wifi-status"
	_, err := parser.Parse()

	if err != nil {
		log.Fatalf("option not loaded. %v", err)
	}

	switch {
	case cmdOption.Version:
		fmt.Printf("%s \n", version)
	default:
		run(cmdOption.GetFilename(), cmdOption.Log, cmdOption.GetTime())
	}

}

func run(filename string, logFlag bool, time int) {
	config, err := config.NewConfig(filename)
	if err != nil {
		log.Fatalf("config not loaded. %v", err)
	}
	config.LogOutput = logFlag
	config.Time = time
	app.Run(config)
}
