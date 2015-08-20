# tgbot
Go library as a wrapper of the Telegram Bot API.

## Example usage

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

## Development

 * `git clone` this repository or `go get github.com/nubunto/tgbot`
 * run `go test -token=<your-bot-token>` to test with a bot of yours.
 * happy hacking

**Please note that the API is still under heavy development.** It still isn't ready to use.