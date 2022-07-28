package client

import (
	"time"

	nilbot "github.com/Ink-33/NilBot"
	"github.com/Ink-33/NilBot/driver"
)

// Bot defines the onebot client.
type Bot struct {
	APICaller   nilbot.APICaller
	SelfID      string
	AccessToken string
	Timeout     uint8
	URL         string
}

// Connect to onebot server.
func (b *Bot) Connect(handler ...nilbot.EventHandler) {
	if b.Timeout == 0 {
		b.Timeout = 3
	}
	ws := driver.NewWebSocketClient(b.URL, b.AccessToken, time.Duration(b.Timeout)*time.Second)
	ws.Connect()
	go ws.Listen(handler...)
	func() {
		resp, _ := ws.CallAPI(&nilbot.APIRequest{
			Action: "get_login_info",
			Params: nil,
		})
		b.SelfID = resp.Data.Get("user_id").String()
	}()
	b.APICaller = ws
}
