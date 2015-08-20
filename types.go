package tgbot

type TelegramUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type TelegramMessage struct {
	ID   int          `json:"message_id"`
	From TelegramUser `json:"from"`
	Text string       `json:"text"`
}

type TelegramUpdate struct {
	ID      int             `json:"update_id"`
	Message TelegramMessage `json:"message"`
}

type telegramResponse struct {
	Ok bool `json:"ok"`
}

type telegramUpdateResponse struct {
	telegramResponse
	Result []TelegramUpdate `json:"result"`
}

type telegramUserResponse struct {
	telegramResponse
	Result TelegramUser `json:"result"`
}