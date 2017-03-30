package linenotify

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// Service ...
type Service struct {
}

// SendPush ...
func (ln *Service) SendPush(token, text string) error {
	params := url.Values{}
	params.Set("message", text)
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", body)

	// Headers
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Fetch Request
	_, err = client.Do(req)

	return err
}
