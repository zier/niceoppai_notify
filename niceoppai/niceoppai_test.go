package niceoppai

import "testing"

func TestGetAllCartoonDetail(t *testing.T) {
	n := &Service{}
	_, err := n.GetAllCartoonDetail()

	if err != nil {
		t.Error(err)
	}
}
