package sms

type SMS interface {
	Send(body string) error
}
