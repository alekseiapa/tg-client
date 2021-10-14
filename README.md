# Micha

[![Build Status](https://travis-ci.org/onrik/micha.svg?branch=master)](https://travis-ci.org/onrik/micha)
[![Coverage Status](https://coveralls.io/repos/github/onrik/micha/badge.svg?branch=master)](https://coveralls.io/github/onrik/micha?branch=master)
[![GoDoc](https://godoc.org/github.com/alekseiapa/tg-client?status.svg)](https://godoc.org/github.com/alekseiapa/tg-client)
[![Gitter](https://badges.gitter.im/onrik/micha.svg)](https://gitter.im/onrik/micha)

Client lib for [Telegram bot api](https://core.telegram.org/bots/api). Supports **Bot API v2.3** (of 21th Nov 2016).

##### Simple echo bot example:
```go
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

```
