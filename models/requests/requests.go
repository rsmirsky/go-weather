package req

type NewHistory struct {
	TelegramChatID   int64
	TelegramUserName string
	Command          string
	IsCity           bool
}
