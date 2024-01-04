package main

import (
	"flag"
	"time"
)

func main() {
	var config_path string
	var user string
	var host string
	var ip string

	flag.StringVar(&config_path, "config_path", "config.json", "Path to config.json")
	flag.StringVar(&user, "user", "", "")
	flag.StringVar(&host, "host", "", "")
	flag.StringVar(&ip, "ip", "", "")
	flag.Parse()

	config := &Config{}

	cfgunm := config.Unmarshal(config_path)

	if cfgunm != nil {
		panic(cfgunm)
	}

	tgbot := &TgBot{
		Config: config,
	}

	if tgbot.Create() != nil {
		panic(tgbot.Create())
	}

	messageTemplate := &MessageTemplate{}

	var msg string = messageTemplate.Create(
		host,
		user,
		time.Now(),
		ip,
	)

	tgbot.Send(msg)
}
