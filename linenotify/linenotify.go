package linenotify

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Service ...
type Service struct {
}

// SendPush ...
func (ln *Service) SendPush(token, text string, thumbnail string) error {
	params := url.Values{}
	params.Set("message", text)
	params.Set("imageThumbnail", thumbnail)
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", body)

	// Headers
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 401 {
		return errors.New("invalid token")
	}

	return nil
}
