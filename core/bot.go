package core

// BotClient is the interface of bot instance
type BotClient interface {
	Send() error
	Regist()
}