package tgbot

import (
	"net/http"
	"encoding/json"
	"fmt"
)

const telegramAPI = "https://api.telegram.org/bot%s/%s"

type TelegramBot struct {
	token string // the bot token of the API
}

func NewBot(token string) *TelegramBot {
	return &TelegramBot{token}
}

func (bot TelegramBot) GetUpdates() ([]TelegramUpdate, error) {
	resp, err := http.Get(fmt.Sprintf(telegramAPI, bot.token, "getUpdates"))
	if err != nil {
		return []TelegramUpdate{}, err
	}
	defer resp.Body.Close()
	var updates telegramUpdateResponse
	if err := json.NewDecoder(resp.Body).Decode(&updates); err != nil {
		return []TelegramUpdate{}, err
	}
	return updates.Result, nil
}

func (bot TelegramBot) GetMe() (TelegramUser, error) {
	resp, err := http.Get(fmt.Sprintf(telegramAPI, bot.token, "getMe"))
	if err != nil {
		return TelegramUser{}, err
	}
	defer resp.Body.Close()
	var retBot telegramUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&retBot); err != nil {
		return TelegramUser{}, err
	}
	return retBot.Result, nil
}