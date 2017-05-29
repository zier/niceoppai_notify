package linenotify

import "testing"

func TestSendPush(t *testing.T) {
	n := &Service{}
	err := n.SendPush("ZzdbtQb4Zs4zf3XCxJD2QnlCvt5wWBNE9DRm9JY8UY4", "test", "http://www.google.com")

	if err != nil {
		t.Error(err)
	}
}

func TestSendPushInvalid(t *testing.T) {
	n := &Service{}
	err := n.SendPush("invalid_token", "test", "http://www.google.com")

	if err.Error() != "invalid token" {
		t.Error("token should invalid")
	}
}
