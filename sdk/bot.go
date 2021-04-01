package nilbot

type BotClient interface {
	Send() error
	Regist()
}