# Micha

[![Tests](https://github.com/alekseiapa/tg-client/workflows/Tests/badge.svg)](https://github.com/alekseiapa/tg-client/actions)
[![Coverage Status](https://coveralls.io/repos/github/onrik/micha/badge.svg?branch=master)](https://coveralls.io/github/onrik/micha?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alekseiapa/tg-client)](https://goreportcard.com/report/github.com/alekseiapa/tg-client)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/alekseiapa/tg-client)](https://pkg.go.dev/github.com/alekseiapa/tg-client)
[![Gitter](https://badges.gitter.im/onrik/micha.svg)](https://gitter.im/onrik/micha)

Client lib for [Telegram bot api](https://core.telegram.org/bots/api). Supports **Bot API v2.3.1** (of 4th Dec 2016).

##### Simple echo bot example:
```go
package main

import (
    "log"
	
    "github.com/alekseiapa/tg-client"
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
