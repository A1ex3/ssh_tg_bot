package main

import (
	"flag"
	"time"
)

func main() {
	var config_path string
	flag.StringVar(&config_path, "config_path", "config.json", "Path to config.json")
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

	sshInfo := &SshInfo{}
	sshInfoGet := sshInfo.Get() //

	if sshInfoGet != nil {
		panic(sshInfoGet)
	}

	var msg string = messageTemplate.Create(
		sshInfo.HostIp,
		time.Now(),
		sshInfo.RemoteIp,
		sshInfo.RemoteIpPort,
	)

	tgbot.Send(msg)
}
