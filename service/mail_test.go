package service

import "testing"

func TestEmailServiceSend(t *testing.T) {
	service, err := NewMailService()
	if err != nil {
		t.Fatal(err)
	}
	err = service.Send("hello", "this is a testing mail", []string{"me@forseason.vip"})
	if err != nil {
		t.Fatal(err)
	}
}
