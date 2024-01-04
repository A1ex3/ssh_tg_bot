package main

import (
	"log"
	"sync"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct{
	Bot *tgbotapi.BotAPI
	Config *Config
}

func (tgBot *TgBot) Create() error {
	bot, err := tgbotapi.NewBotAPI(tgBot.Config.TgBotApiToken)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	bot.Debug = tgBot.Config.TgBotDebug

	tgBot.Bot = bot
	return nil
}

func (tgBot *TgBot) Send(message string){
	var userListLen int = len(tgBot.Config.UserList)

	var wg sync.WaitGroup

	wg.Add(userListLen)
	for _, i := range tgBot.Config.UserList{
		go func(i int64){
			defer wg.Done()
			tgBot.Bot.Send(tgbotapi.NewMessage(i, message))
		}(i)
	}
	wg.Wait()
}