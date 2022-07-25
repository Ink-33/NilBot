package utils

import (
	"io"
	"net/http"
	"time"
	"unsafe"
)

// Get web Content by using GET request.
func Get(addr string) (body []byte, err error) {
	var content []byte
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	request, err := http.NewRequest("GET", addr, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4239.0 Safari/537.36")
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// PostJSON to addr and get response body
func PostJSON(addr string, postbody io.Reader) (body []byte, err error) {
	var content []byte
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	request, err := http.NewRequest("POST", addr, postbody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4117.2 Safari/537.36")
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	content, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// B2S convert byte slice to string.
func B2S(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
