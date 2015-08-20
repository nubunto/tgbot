# tgbot
Go library as a wrapper of the Telegram Bot API.

## Example

	package main
	
	import (
		"github.com/nubunto/tgbot"
		"fmt"
	)

	func main() {
		bot := tgbot.NewBot("<your-bot-token>")
		myBot, err := bot.GetMe()

		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
		fmt.Printf("Hello, my name is %s", myBot.Username)
	}

**Please note that the API is still under heavy development.** It still isn't ready to use.ill isn't ready to use.