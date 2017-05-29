package linenotify

import "testing"

func TestSendPush(t *testing.T) {
	n := &Service{}
	err := n.SendPush("2PxNQT9aW7DWT16PmJyctLrw31VptTNs4mOtYsSeIqa", "test", "http://www.niceoppai.net/wp-content/manga/cover/tbn/talesofdemonandgod_200x0.png")

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
