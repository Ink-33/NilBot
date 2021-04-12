package core

type BotClient interface {
	Send() error
	Regist()
}