package driver

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/RomiChan/websocket"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
	"golang.org/x/sync/syncmap"

	nilbot "github.com/Ink-33/NilBot"
	"github.com/Ink-33/NilBot/internal/logger"
)

// WSClient defines the default websocket client.
type WSClient struct {
	AccessToken string
	conn        *websocket.Conn
	lock        sync.Mutex
	rayIDMap    syncmap.Map
	Timeout     time.Duration
	URL         string
}

// NewWebSocketClient returns a websocket client.
func NewWebSocketClient(url, accessToken string, timeout time.Duration) *WSClient {
	return &WSClient{URL: url, AccessToken: accessToken, Timeout: timeout}
}

// Connect websocket server.
func (c *WSClient) Connect() {
	logger.Info("Try connecting to websocket server %v", c.URL)
	header := http.Header{
		"User-Agent": []string{"NilBot/dev"},
	}
	if c.AccessToken != "" {
		header["Authorization"] = []string{"Bear" + " " + c.AccessToken}
	}

TRY:
	conn, resp, err := websocket.DefaultDialer.Dial(c.URL, header)
	if err != nil {
		logger.Warn("An error occured while trying connecting to websocket server %v : %v", c.URL, err.Error())
		time.Sleep(3 * time.Second)
		goto TRY
	}
	c.conn = conn
	resp.Body.Close()

	logger.Info("Connected to websocket server %v .", c.URL)
}

// Listen websocket events.
func (c *WSClient) Listen(handler ...nilbot.EventHandler) {
	for {
		mtype, payload, err := c.conn.ReadMessage()
		if err != nil {
			logger.Warn("Disconnected from websocket server %v", c.URL)
			c.Connect()
		}
		if mtype == websocket.TextMessage {
			resp := gjson.ParseBytes(payload)
			if resp.Get("echo").Exists() {
				if ch, ok := c.rayIDMap.LoadAndDelete(resp.Get("echo").String()); ok {
					ch.(chan *nilbot.APIResponse) <- &nilbot.APIResponse{
						Data:    resp.Get("data"),
						Echo:    resp.Get("echo").Str,
						Msg:     resp.Get("msg").Str,
						RetCode: resp.Get("retcode").Int(),
						Status:  resp.Get("status").Str,
						Wording: resp.Get("wording").Str,
					}
					close(ch.(chan *nilbot.APIResponse))
				}
			} else {
				for k := range handler {
					go handler[k].Handle(payload, c)
				}
			}
		}
	}
}

// CallAPI sends websocket request.
func (c *WSClient) CallAPI(req *nilbot.APIRequest) (resp *nilbot.APIResponse, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	req.Echo = uuid.NewString()
	ch := make(chan *nilbot.APIResponse, 1)
	c.rayIDMap.Store(req.Echo, ch)
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	c.lock.Lock()
	err = c.conn.WriteMessage(websocket.TextMessage, payload)
	c.lock.Unlock()
	if err != nil {
		logger.Warn("An error occured while requesting API : %v\n", err.Error())
		return nil, err
	}

	select {
	case resp, ok := <-ch:
		if !ok {
			return nil, errors.New("channel closed")
		}
		return resp, nil
	case <-time.After(c.Timeout):
		return nil, errors.New("time out")
	}
}
