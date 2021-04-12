package core

import "sync"

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

func (e *ParsedMessageElement) Get(key string) string {
	r := e.data["data"].(map[string]interface{})[key]
	if r != nil {
		return r.(string)
	}
	return ""
}

func (e *ParsedMessageElement) GetElement() MsgElement {
	return e.data
}

func (e *ParsedMessageElement) Type() string {
	return e.data["type"].(string)
}
