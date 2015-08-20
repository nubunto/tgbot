package main

import (
	"github.com/nubunto/tgbot"
	"fmt"
)

func main() {
	bot := tgbot.NewBot("<your-bot-token>")
	user, _ := bot.GetMe()
	fmt.Printf("hello, my name is %s", user.Username)
}