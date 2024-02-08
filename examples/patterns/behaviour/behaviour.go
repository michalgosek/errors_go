package behaviour

import "time"

type timeoutErr string

func (t timeoutErr) Error() string { return string(t) }

func (t timeoutErr) IsTimeout() bool { return true }

func Connection() error {
	time.Sleep(2 * time.Second)
	return timeoutErr("connection failed")
}
