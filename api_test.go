package tgbot

import (
	"testing"
	"os"
	"flag"
)

var (
	token = flag.String("token", "the bot token to test against", "")
)

var getMeTests = []string {
	"icantevenbot",
}


func TestMain(m *testing.M) {
	flag.Parse()
	code := m.Run()
	os.Exit(code)
}

func TestGetMe(t *testing.T) {
	bot := NewBot(*token)
	for i, test := range getMeTests {
		if user, err := bot.GetMe(); err != nil {
			t.Errorf("Got unexpected error at %d. Error: %s", i, err.Error())
		} else {
			if test != user.Username {
				t.Errorf("Unexpected bot username; Expected [%s], got [%s]", test, user.Username)
			}
		}
	}
}