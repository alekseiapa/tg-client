# Telegram Bot API Client 

Go client library for interacting with the [Telegram Bot API](https://core.telegram.org/bots/api). It allows developers to easily create Telegram bots with Go. Whether you want to create a simple echo bot or a complex bot that uses all available API features.

---

## 📦 Installation

Ensure you have [Go](https://golang.org) installed and configured. To install the library, use the following command:

```bash
go get github.com/alekseiapa/tg-client
```

---

## ⚡️ Quick Start

Here's a quick example to create a simple echo bot:

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

### Custom API Server

It's possible to use Client with a custom API server:

```go
package main

import (
    "log"

    "github.com/alekseiapa/tg-client"
)

func main() {
    bot, err := tg.NewBot(
        "<token>",
        tg.WithAPIServer("http://127.0.0.1:8081"),
    )
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

---

## 🌟 Features

- **Comprehensive API Coverage**: Support for all Telegram Bot API methods.
- **Customizable**: Use custom API servers and HTTP clients.
- **Streaming**: Long polling with `Start()` and `Stop()` methods to manage updates efficiently.
- **Webhook Support**: Easily configure webhooks.
- **Error Handling**: Detailed error messages for efficient debugging.
- **Middleware**: Chainable middleware.

---

## 📚 Documentation

- **Telegram Bot API Documentation**: Visit the [official documentation](https://core.telegram.org/bots/api)


---

## 🧪 Testing

Run the test suite using:

```bash
go test -covermode=atomic -coverprofile=profile.cov ./...
```

---


### Credits & Acknowledgements
This project was inspired by the original project by Micha.

All rights and credits for the original idea, design, and concept belong to the original author. I used this project solely for learning and study purposes to improve my Go programming skills.

Thank you to the original author!
