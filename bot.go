package nilbot

import "github.com/tidwall/gjson"

// APICaller defines basic abilities of supported api caller.
type APICaller interface {
	CallAPI(req *APIRequest) (resp *APIResponse, err error)
}

// EventHandler basic abilities of event handler.
type EventHandler interface {
	Handle(payload []byte, apiCaller APICaller)
}

// APIRequest is used to communicate with onebot client.
// https://github.com/botuniverse/onebot-11/blob/master/communication/ws.md
type APIRequest struct {
	Action string         `json:"action"`
	Echo   string         `json:"echo"`
	Params map[string]any `json:"params"`
}

// APIResponse is the response of calling API
// https://github.com/botuniverse/onebot-11/blob/master/communication/ws.md
type APIResponse struct {
	Data    gjson.Result `json:"data"`
	Echo    string       `json:"echo"`
	Msg     string       `json:"msg"`
	RetCode int64        `json:"retcode"`
	Status  string       `json:"status"`
	Wording string       `json:"wording"`
}
