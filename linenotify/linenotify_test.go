package linenotify

import "testing"

func TestSendPush(t *testing.T) {
	n := &Service{}
	err := n.SendPush("TOKEN_TEST", "test")

	if err != nil {
		t.Error(err)
	}
}
