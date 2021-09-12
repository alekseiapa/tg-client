package main

import (
	"github.com/alekseiapa/tg-client"
	"log"
)

func main() {
	bot, err := tg.NewBot("<token>")
	if err != nil {
		log.Fatal(err)
	}

	go bot.Start()

	for update := range bot.Updates() {
		if update.Message != nil {
			bot.SendMessage(update.Message.Chat.ID, update.Message.Text, nil)
		}
	}
}
