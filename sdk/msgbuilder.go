package nilbot

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

// MsgElement is an element of array message.
type MsgElement map[string]interface{}

// MsgArray contains MsgElements, used to send and revice message by using OneBot array message format.
type MsgArray []MsgElement

// MsgBuilder is a builder used to build array messgae.
type MsgBuilder struct {
	msgMap MsgArray
}

// MsgBuilderArg is used to represent complex parameters in the message element.
// Usually if the function that can recive more than one parameter need to use this.
//
// For more detail, please read document.
type MsgBuilderArg interface{}

// At append an at message element to message array.
func (m *MsgBuilder) At(qq string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "at", "data": MsgElement{"qq": qq}})
	return m
}

// CardImage append a CardImage message element to message array.
func (m *MsgBuilder) CardImage(args MsgBuilderArg) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "cardimage", "data": args})
	return m
}

// Face append a face message element to message array.
func (m *MsgBuilder) Face(id string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "face", "data": MsgElement{"id": id}})
	return m
}

// Forward append a forward message element to message array.
func (m *MsgBuilder) Forward(id string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "forward", "data": MsgElement{"id": id}})
	return m
}

// Gife append an gift message element to message array.
func (m *MsgBuilder) Gift(qq, giftID string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "gift", "data": MsgElement{"qq": qq, "id": giftID}})
	return m
}

// Image append an image message element to message array.
func (m *MsgBuilder) Image(args MsgBuilderArg) *MsgBuilder {
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

// JSON append a JSON message element to message array.
func (m *MsgBuilder) JSON(json string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "json", "data": MsgElement{"data": json}})
	return m
}

// Music append a music message element to message array.
func (m *MsgBuilder) Music(args MsgBuilderArg) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "music", "data": args})
	return m
}

// Node append a forward node element to message array.
//
// You should read document before using this.
func (m *MsgBuilder) Node(args MsgBuilderArg) *MsgBuilder {
	switch args.(type) {
	case MsgElement:
		m.msgMap = append(m.msgMap, MsgElement{"type": "node", "data": args})
	case string:
		m.msgMap = append(m.msgMap, MsgElement{"type": "node", "data": MsgElement{"id": args}})
	default:
		m.msgMap = append(m.msgMap, MsgElement{"type": "text", "data": MsgElement{"text": fmt.Sprintf("Unexpected args type: %T", args)}})
	}
	return m
}

// Poke append an Poke message element to message array.
func (m *MsgBuilder) Poke(qq string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "poke", "data": MsgElement{"qq": qq}})
	return m
}

// Record append a record message element to message array.
func (m *MsgBuilder) Record(args MsgBuilderArg) *MsgBuilder {
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
func (m *MsgBuilder) Share(args MsgBuilderArg) *MsgBuilder {
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
func (m *MsgBuilder) Vedio(args MsgBuilderArg) *MsgBuilder {
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

// XML append a XML message element to message array.
func (m *MsgBuilder) XML(xml string) *MsgBuilder {
	m.msgMap = append(m.msgMap, MsgElement{"type": "xml", "data": MsgElement{"data": xml}})
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
