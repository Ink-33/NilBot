package nilbot

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

// MsgElement is an element of array message.
type MsgElement map[string]interface{}

// MsgBuilder is a builder used to build array messgae.
type MsgBuilder struct {
	msgMap []interface{}
}

// At append an at message element to message array.
func (m *MsgBuilder) At(qq string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "at", "data": MsgElement{"qq": qq}})
	return m
}

// Face append a face message element to message array.
func (m *MsgBuilder) Face(id string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "face", "data": MsgElement{"id": id}})
	return m
}

// Image append an image message element to message array.
func (m *MsgBuilder) Image(args interface{}) *MsgBuilder {
	switch args.(type) {
	case MsgElement:
		m.msgMap = append(m.msgMap, MsgElement{"type": "image", "data": args})
	case string:
		m.msgMap = append(m.msgMap, MsgElement{"type": "image", "data": MsgElement{"file": args}})
	default:
		m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": fmt.Sprintf("Unexpected args type: %T", args)}})
	}
	return m
}

// Music append a music message element to message array.
func (m *MsgBuilder) Music(args interface{}) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "music", "data": args})
	return m
}

// Record append a record message element to message array.
func (m *MsgBuilder) Record(args interface{}) *MsgBuilder {
	switch args.(type) {
	case MsgElement:
		m.msgMap = append(m.msgMap, MsgElement{"type": "record", "data": args})
	case string:
		m.msgMap = append(m.msgMap, MsgElement{"type": "record", "data": MsgElement{"file": args}})
	default:
		m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": fmt.Sprintf("Unexpected args type: %T", args)}})
	}
	return m
}

// Share append an share message element to message array.
func (m *MsgBuilder) Share(args interface{}) *MsgBuilder {
	switch args.(type) {
	case MsgElement:
		m.msgMap = append(m.msgMap, MsgElement{"type": "share", "data": args})
	case string:
		m.msgMap = append(m.msgMap, MsgElement{"type": "share", "data": MsgElement{"url": args}})
	default:
		m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": fmt.Sprintf("Unexpected args type: %T", args)}})
	}
	return m
}

// Text append a text message element to message array.
func (m *MsgBuilder) Text(text string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": text}})
	return m
}

// Vedio append a vedio message element to message array.
func (m *MsgBuilder) Vedio(args interface{}) *MsgBuilder {
	switch args.(type) {
	case MsgElement:
		m.msgMap = append(m.msgMap, MsgElement{"type": "vedio", "data": args})
	case string:
		m.msgMap = append(m.msgMap, MsgElement{"type": "vedio", "data": MsgElement{"file": args}})
	default:
		m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": fmt.Sprintf("Unexpected args type: %T", args)}})
	}
	return m
}

// Build marshal the message array to []byte.
func (m *MsgBuilder) Build() ([]byte, error) {
	b, err := json.Marshal(m.msgMap)
	if err != nil {
		return nil, err
	}
	return b, nil
}
