package message

import (
	"encoding/json"
	"sync"
)

// ParsedMessageElement is
type ParsedMessageElement struct {
	data MsgElement
}

// MsgParser is used to parser message of array format.
func MsgParser(message []byte) ([]ParsedMessageElement, error) {
	var pm = &[]MsgElement{}
	err := json.Unmarshal(message, pm)
	if err != nil {
		return nil, err
	}
	l := len(*pm)
	var pme = make([]ParsedMessageElement, l)
	wg := sync.WaitGroup{}
	wg.Add(l)
	for i := 0; i < l; i++ {
		go func(i int, pm []MsgElement) {
			pme[i].data = pm[i]
			wg.Done()
		}(i, *pm)
	}
	wg.Wait()
	return pme, nil
}

// Get value form ParsedMessageElement
func (e *ParsedMessageElement) Get(key string) any {
	return e.data["data"].(map[string]any)[key]
}

// GetElement form ParsedMessageElement
func (e *ParsedMessageElement) GetElement() MsgElement {
	return e.data
}

// Type reports the type of ParsedMessageElement
func (e *ParsedMessageElement) Type() string {
	return e.data["type"].(string)
}
