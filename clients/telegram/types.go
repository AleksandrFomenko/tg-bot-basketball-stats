package telegram

type UdpatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
	//CallbackQuery *CallbackQuery   `json:"callback_query"`
}

type IncomingMessage struct {
	Text        string      `json:"text"`
	From        From        `json:"from"`
	Chat        Chat        `json:"chat"`
	ReplyMarkup ReplyMarkup `json:"reply_markup"`
}
type CallbackQuery struct {
	Id              string `json:"id"`
	InlineMessageId string `json:"inline_message_id"`
}
type From struct {
	Username string `json:"username"`
}
type Chat struct {
	ID int `json:"id"`
}

type ReplyMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"url"`
}
